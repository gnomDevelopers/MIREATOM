// const spliteSymbol = '|';
// const whiteSymbol = '_';

// // \cfrac{\sqrt[3]{x\int_{a}^{b}{adx}}}{b\log_{10}{\sum_{i=a}^b{n}}}

// <mrow><mstyle displaystyle="true" scriptlevel="0"><mfrac><mroot><mrow><mi>x</mi><msubsup><mo>∫</mo><mi>a</mi><mi>b</mi></msubsup><mrow><mi>a</mi><mi>d</mi><mi>x</mi></mrow></mrow><mn>3</mn></mroot><mrow><mi>b</mi><msub><mrow><mi>log</mi><mo>⁡</mo></mrow><mn>10</mn></msub><mrow><msubsup><mo>∑</mo><mrow><mi>i</mi><mo>=</mo><mi>a</mi></mrow><mi>b</mi></msubsup><mi>n</mi></mrow></mrow></mfrac></mstyle></mrow>

// export function parseLatexFromHTML(html: HTMLElement){
//   const textList = [] as string[];
//   getTextContent(html, textList);
//   console.log(textList);

//   //заменяем пустые поля на спец символ
//   textList.forEach((item, index) => {
//     if(item === '') textList[index] = whiteSymbol;
//   });

//   let formula = textList[textList.length - 1];
//   let text = textList.splice(0, textList.length - 1).join(spliteSymbol);
//   let cursorTextPos = 0;
//   let cursorFormulaPos = 0;

//   console.log('formula: ', formula);
//   console.log('text: ', text);

//   while(cursorTextPos < text.length){

//     //---сдвигаем курсор в формуле---

//     //если в формуле начинается запись функции пропускаем ее инициализацию
//     if(['\\'].includes(formula[cursorFormulaPos])){
//       // сдвигаем курсор право, пока не встретим аргумент{}
//       while(!['{'].includes(formula[cursorFormulaPos])){
//         cursorFormulaPos++;
//       }
//       cursorFormulaPos++;
//     }

//     //если в строке разделительный символ
//     if(isSpecialSymbol(text, spliteSymbol, cursorTextPos)){
//       //если нашли рофло символ, двигаемся вперед пока они не пропадут
//       while(['{', '}', ' '].includes(formula[cursorFormulaPos])){
//         cursorFormulaPos++;
//       }
//     }
    
//     //---сдвигаем курсор в текстовой строке---

//     //если встречаем разделительный символ - пропускаем его
//     if(isSpecialSymbol(text, spliteSymbol, cursorTextPos)){
//       cursorTextPos += spliteSymbol.length;
//     }

//     //---теперь все курсоры находятся на тексте---

//     //если символы в формуле и в тексте совпали - двигаем оба курсора дальше
//     if(text[cursorTextPos] === formula[cursorFormulaPos]){
//       cursorTextPos++;
//       cursorFormulaPos++;
//       continue;
//     }

//     //добавляем в формулу несовпадающий символ из текста
//     if(isSpecialSymbol(text, whiteSymbol, cursorTextPos)){ // если это пустой символ ничего не добавляем
//       // запоминаем начало формулы
//       const formulaBegin = formula.slice(0, cursorFormulaPos);
//       //пропускаем все аргументы пока не дойдем до закрывающего символа
//       while(!['}'].includes(formula[cursorFormulaPos]) && cursorFormulaPos < formula.length){
//         cursorFormulaPos++;
//       }
//       //обновляем формулу
//       formula = formulaBegin + '' + formula.slice(cursorFormulaPos, formula.length);
//     }
//     else { // иначе добавляем отсутствующий символ
//       // if(cursorTextPos === text.length - 1 && )
//       formula = formula.slice(0, cursorFormulaPos) + text[cursorTextPos] + formula.slice(cursorFormulaPos, formula.length);
//     }

//     //сдвигаем курсоры дальше
//     cursorTextPos++;
//     cursorFormulaPos++;
//   }

//   console.log('new formula: ', formula);
//   return formula;
// }

// function getTextContent(node: Element, textList: string[]){
//   if(node.children.length === 0){
//     if(node.textContent !== null) {
//       textList.push(node.textContent);
//     }
//   }
//   else{
//     for(let child of node.children){
//       getTextContent(child, textList);
//     }
//   }
// }


// //проверка текущей подстроки на разделительный символ
// function isSpecialSymbol(string: string, specialSymbol: string, pos: number){
//   return (pos <= string.length - specialSymbol.length && string.slice(pos, pos + specialSymbol.length) === specialSymbol);
// }


// \cfrac{\sqrt[3]{x\int_{a}^{b}{adx}}}{b\log_{10}{\sum_{i=a}^b{n}}}

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


function latexFromHTML(html: string){
  let latex = "";

  //Basic handling of fractions
  const fractionMatch = html.match(/<mfrac><mi>(.*?)<\/mi><mi>(.*?)<\/mi><\/mfrac>/);
  if (fractionMatch) {
    console.log('fractionMatch: ', fractionMatch);
    latex += "\\cfrac{" + latexFromHTML(fractionMatch[1]) + "}{" + latexFromHTML(fractionMatch[2]) + "}";
  }

  //Basic handling of square roots
  const sqrtMatch = html.match(/<msqrt>(.*?)<\/msqrt>/);
  if (sqrtMatch) {
    latex += "\\sqrt{" + latexFromHTML(sqrtMatch[1]) + "}";
  }

  //<msubsup><mo>∫</mo><mi>a</mi><mi>b</mi></msubsup><mrow><mi>a</mi><mi>d</mi><mi>x</mi></mrow>
  //Basic handling of integrals
  const integralMatch = html.match(/<msubsup><mo>∫<\/mo><mi>(.*?)<\/mi><mi>(.*?)<\/mi><sub>(.*?)<\/sub><sup>(.*?)<\/sup>(.*?)/);
  if (integralMatch) {
    latex += "\\int_{" + latexFromHTML(integralMatch[1]) + "}^{" + latexFromHTML(integralMatch[2]) + "}{" + latexFromHTML(integralMatch[3]) + "}";
  }


  //Basic handling of logs
  const logMatch = html.match(/<msub><mi>log<\/mi><sub>(.*?)<\/sub>(.*?)<\/msub>/);
  if (logMatch) {
    latex += "\\log_{" + latexFromHTML(logMatch[1]) + "}{" + latexFromHTML(logMatch[2]) + "}";
  }


  //Handle remaining text, removing HTML tags
  return latex === '' ? html : latex;
}


export function parseLatexFromHTML(parentElement: Element) {
  const mrow = getTextContent(parentElement);
  const html = mrow?.innerHTML || '';

  let formula = latexFromHTML(html);
  console.log(formula);
  return formula;
}



// function latexFromHTML(html) {
//   let latex = "";

//   // Use a more robust regular expression to handle various fraction formats.
//   const fractionMatch = html.match(/<mfrac>((?:.|\n)*?)<\/mfrac>/i);

//   if (fractionMatch) {
//     console.log(fractionMatch);
//     const numerator = fractionMatch[1].trim();
//     const denominator = fractionMatch[2].trim(); // Fixed, using the first part only
//     latex +=
//       "\\cfrac{" +
//       extractContent(numerator) +
//       "}{" +
//       extractContent(denominator) +
//       "}";
//   } 
//   else {
//     return html; // Return the input if no fraction is found
//   }
//   return latex;
// }


// //Helper function to correctly extract content
// function extractContent(html) {
//   //Regular expressions for different elements:
//   const miMatch = html.match(/<mi>(.*?)<\/mi>/i);
//   const mnMatch = html.match(/<mn>(.*?)<\/mn>/i);
//   const innerMatch = html.match(/<mrow>((?:.|[\r\n])*?)<\/mrow>/i);
  
//   let extractedContent = "";


//   if (miMatch) {
//     extractedContent = miMatch[1].trim(); //trim() to remove possible extra whitespace.
//   } else if (mnMatch) {
//     extractedContent = mnMatch[1].trim();
//   } else if (innerMatch) { // Recursive call for nested expressions within <mrow>
//     extractedContent = extractContent(innerMatch[1]);
//   } else {
//     extractedContent = html; // Handle cases with no match or other types of contents
//   }

//   return extractedContent;

// }



// // Example usage (test cases):

// const html1 = "<mfrac><mn>2</mn><mi>b</mi></mfrac>";
// const html2 = "<mfrac><mi>a</mi><mn>2</mn></mfrac>";
// const html3 = "<mfrac><mrow><mn>2</mn><mi>a</mi></mrow><mi>b</mi></mfrac>";
// const html4 = "<mfrac><mi>a</mi><mrow><mn>2</mn><mi>b</mi></mrow></mfrac>";
// const html5 = "<mfrac><mrow><mn>2</mn><mi>a</mi><mo>+</mo><mn>1</mn></mrow><mrow><mn>2</mn><mi>b</mi><mo>+</mo><mn>3</mn></mrow></mfrac>"; // Example with additional content
// const html6 = "<mfrac><mi>a</mi><mi>b</mi></mfrac>"; // Original example
// const html7 = "<mfrac><mrow><mi>x</mi><mo>+</mo><mn>1</mn></mrow><mn>2</mn></mfrac>"; // Another example with nested expressions
// const html8 = "<mfrac><mi>a</mi><mrow><mi>b</mi><mo>+</mo><mn>1</mn></mrow></mfrac>";


// console.log(latexFromHTML(html1)); // Output: \cfrac{2}{b}
// console.log(latexFromHTML(html2)); // Output: \cfrac{a}{2}
// console.log(latexFromHTML(html3)); // Output: \cfrac{2a}{b}
// console.log(latexFromHTML(html4)); // Output: \cfrac{a}{2b}
// console.log(latexFromHTML(html5)); // Output: \cfrac{2a + 1}{2b + 3}
// console.log(latexFromHTML(html6)); // Output: \cfrac{a}{b}
// console.log(latexFromHTML(html7)); // Output: \cfrac{x + 1}{2}
// console.log(latexFromHTML(html8)); // Output: \cfrac{a}{b + 1}