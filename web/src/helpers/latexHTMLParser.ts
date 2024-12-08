export function parseLatexFromHTML(html: HTMLElement){
  const textList = [] as string[];
  getTextContent(html, textList);
  console.log(textList);
}

function getTextContent(node: Element, textList: string[]){
  if(node.children.length === 0){
    if(node.textContent !== null) {
      console.log('textList before: ', textList);
      console.log('new text: ', node.textContent);
      textList.push(node.textContent);
      console.log('textList after: ', textList);
    }
  }
  else{
    for(let child of node.children){
      getTextContent(child, textList);
    }
  }
}