// \cfrac{\sqrt[3]{x\int_{a}^{b}{adx}}}{b\log_{10}{\sum_{i=a}^{b}{n}}}
// \cfrac{\sqrt[3]{x\int_{a}^{b}{adx}}}{b\log_{10}{\sum_{i=0}^{100}{n}}}\lim_{a \to b}{15x^2}

import katex from "katex";

function getTextContent(node: Element){
  if(node.tagName === 'mrow'){
    return node;
  }
  else{
    for(let child of node.children){
      return getTextContent(child);
    }
  }
}


function latexFromHTML(html: Element){
  let latex = "";

  console.log('element: ', html);

  //если это текстовый элемент то возвращаем его содержимое
  if(['mi', 'mo', 'mn'].includes(html.tagName)){
    if(html.textContent === null) return '';
    if(['→'].includes(html.textContent)) return ' \\to ';
    return html.textContent;
  }

  //иначе смотрим тип элемента и его детей

  //если это дробь
  if(html.tagName === 'mfrac'){
    latex += `\\cfrac{${latexFromHTML(html.children[0])}}{${latexFromHTML(html.children[1])}}`;
    return latex;
  }
  //если это корень с установленной степенью
  if(html.tagName === 'mroot'){
    latex += `\\sqrt[{${latexFromHTML(html.children[1])}}]{${latexFromHTML(html.children[0])}}`;
    return latex;
  }
  //если это квадратный корень
  if(html.tagName === 'msqrt'){
    latex += `\\sqrt{${latexFromHTML(html.children[0])}}`;
    return latex;
  }
  //если это функция с коэфами вверху и внизу
  if(html.tagName === 'msubsup'){
    switch(html.children[0].textContent){
      case '∫': latex += `\\int_{${latexFromHTML(html.children[1])}}^{${latexFromHTML(html.children[2])}}`; break;
      case '∑': latex += `\\sum_{${latexFromHTML(html.children[1])}}^{${latexFromHTML(html.children[2])}}`; break;
    }
    
    return latex;
  }
  //если это функция с коэфами внизу
  if(html.tagName === 'msub'){
    latex += `\\${latexFromHTML(html.children[0])}_{${latexFromHTML(html.children[1])}}`;
    return latex;
  }
  //если это функция с коэфами вверху
  if(html.tagName === 'msup'){
    latex += `{${latexFromHTML(html.children[0])}}^{${latexFromHTML(html.children[1])}}`;
    return latex;
  }

  //если это елемент-строка или другой рофлоэлемент
  if(html.children.length !== 0){
    for(let i = 0; i < html.children.length; i++){
      latex += latexFromHTML(html.children[i]);
      if(['msubsup', 'msub'].includes(html.children[i].tagName)){
        latex += `{${latexFromHTML(html.children[i+1])}}`;
        i++;
      }
    }
    return latex;
  }
  else return '';
}


export function parseLatexFromHTML(parentElement: Element) {
  const mrow = getTextContent(parentElement)!;
  let formula = latexFromHTML(mrow);
  console.log(formula);
  return formula;
}


export function insertHTMLBeforeCursor(parentElement: HTMLElement, insertFormula: string){
  //если формула пуста, то просто рендерим внутрь формулу
  if(parentElement.children.length === 0){
    //рендерим формулу
    renderKatex(parentElement, insertFormula);
    //выход
    return;
  }

  //если katex отрендерился но начальный mrow пустой, то рендерим сразу в него
  if(getTextContent(parentElement)?.children.length === 0){
    //удаляем всех потомков корневого элемента
    parentElement.innerHTML = '';
    //рендерим в него формулу
    renderKatex(parentElement, insertFormula);
    //выход
    return;
  }

  //получаем позицию курсора
  const selection = window.getSelection();
  //нода в которой стоит курсор и его позиция в тексте
  const recSearch: {node: Node | null, pos: number | null} = {node: null, pos: null};

  // если курсор есть - ищем ноду в которой он находится и позицию курсора
  if (selection !== null && selection.rangeCount !== 0) {
    //что-то на умном
    const range = selection.getRangeAt(0);
    //получаем текстовую ноду и позицию курсора
    Object.assign(recSearch, recursiveSearch(parentElement, range));
  } 

  //если ниче не надено или курсора нет - ищем последний элемент
  if(recSearch.node === null) {
    function findLastNode(html: Node): Node{
      if(html.childNodes.length !== 0){
        return findLastNode(html.childNodes[html.childNodes.length - 1]);
      }
      else{
        return html;
      }
    }

    recSearch.node = findLastNode(getTextContent(parentElement)!);
    recSearch.pos = -1;
  }
  //получаем родительский элемент и деда 
  const selectedNodeParent = recSearch.node.parentElement;
  if(!selectedNodeParent) return;
  const selectedNodeGrandParent = selectedNodeParent.parentElement;
  if(!selectedNodeGrandParent) return;

  //проверяем, не удалился ли родительский элемент из деда
  let grandHasParent = false;
  for(const child of selectedNodeGrandParent.children) {
    if (child === selectedNodeParent) {
      console.log('child: ', child);
      console.log('parent: ', selectedNodeParent);
      grandHasParent = true; 
      break;
    }
  }
  console.log('ОТЕЦ: ', (grandHasParent ? 'НАЙДЕН' : 'НЕ НАЙДЕН'));
  //если отец не нашелся (паранармальщина) 
  if(!grandHasParent){
    console.log('следующий ебень: ', selectedNodeParent.nextSibling);
  }

  console.log('cursorPos: ', recSearch.pos);
  console.log('selectedNode: ', recSearch.node);
  console.log('selectedNodeParent: ', selectedNodeParent);
  console.log('selectedNodegrandParent: ', selectedNodeGrandParent);
  console.log('selectedNodegrandParentChildren: ', selectedNodeGrandParent.children);

  for(const child of selectedNodeGrandParent.children) {
    console.log('child: ', child);
  }

  //создаем временный элемент и в него рендерим latex формулу
  const element = document.createElement('div');
  //рендерим формулу
  renderKatex(element, insertFormula);
  //получаем основную часть формулы без тонны оберток
  const renderedFormula = getTextContent(element);
  //если ниче нет - странно - выход
  if(!renderedFormula) return;

  //если родительский элемент находится в строке
  if(selectedNodeGrandParent.tagName === 'mrow'){
    //расспаковываем все элементы из mrow и добавляем перед или после родительским элементом
    for(const child of renderedFormula.children){
      selectedNodeGrandParent.insertBefore(child, selectedNodeParent);
    }
    //если нужно было добавить формулу после копии родительского элемента 
    if(recSearch.pos === -1){
      //удаляем копию родительского элемента
      selectedNodeGrandParent.removeChild(selectedNodeParent);
      //и вставляем её перед первым дочерним элементом формулы
      selectedNodeGrandParent.insertBefore(selectedNodeParent, renderedFormula.children[0]);
      //т.о. копия родительского элемента будет стоять перед элементами формулы
    }
  }
  else{
    //если нет деда - беда беда
    if(selectedNodeGrandParent === null) return;
    //создаем элемент строку mrow
    const mrowElement = document.createElement('mrow');
    //делаем глуюокую копию родительского элемента
    const newParentNode = selectedNodeParent.cloneNode(true);
    //добавляем копию родительского элемента в mrow
    mrowElement.appendChild(newParentNode);
    //расспаковываем все элементы из mrow формулы и добавляем перед копией родительского элемента
    for(const child of renderedFormula.children){
      mrowElement.insertBefore(child, newParentNode);
    }
    //если нужно было добавить формулу после копии родительского элемента 
    if(recSearch.pos === -1){
      //удаляем копию родительского элемента
      mrowElement.removeChild(newParentNode);
      //и вставляем её перед первым дочерним элементом формулы
      mrowElement.insertBefore(newParentNode, renderedFormula.children[0]);
      //т.о. копия родительского элемента будет стоять перед элементами формулы
    }
    //добавляем деду mrow перед родительским элементом
    selectedNodeGrandParent.insertBefore(mrowElement, selectedNodeParent);
    //удаляем родительский элемент из деда
    selectedNodeGrandParent.removeChild(selectedNodeParent);
  }
}

// \sqrt{x}

// находим текстовую ноду в которой стоит курсор
function recursiveSearch(node: Node, range: Range): {node: Node | null, pos: number | null} {
  if (node.nodeType === Node.TEXT_NODE && range.isPointInRange(node, range.startOffset)) {
    return {node: node, pos: range.endOffset};
  }
  for (let i = 0; i < node.childNodes.length; i++) {
    const foundNode = recursiveSearch(node.childNodes[i], range);
    if (foundNode) {
      return foundNode;
    }
  }
  return {node: null, pos: null};
}

function renderKatex(element: HTMLElement, formula: string){
  katex.render(formula, element, {
    throwOnError: true,
    displayMode: false,
    output: 'mathml',
    trust: false,
  });
}