//all validation functions
import { type IValidAnswer } from "./constants";

//валидация логина пользователя
export function ValidUserLogin(value: string): IValidAnswer{
  if(value.length < 4) {
    return {value: '', error: 'Слишком короткий логин!'};
  }
  if(value.length > 60) {
    return {value: '', error: 'Слишком длинный логин!'};
  }
  // Проверка на корректные символы: буквы и цифры
  if (value.match(/^[a-zA-Z0-9]+$/) === null) {
    return { value: '', error: 'Логин может содержать только буквы и цифры!' };
}

// Проверка на наличие хотя бы одной буквы
if (value.match(/[a-zA-Z]/) === null) {
    return { value: '', error: 'Логин должен содержать хотя бы одну букву!' };
}
  return {value: value, error: ''};
}

//валидация пароля пользователя
export function ValidUserPassword(value: string): IValidAnswer{
  if(value.match(/[a-zA-Z]+/) === null){
    return {value: '', error: 'Пароль должен содержать латинские буквы в обоих регистрах!'};
  }
  if(value.match(/[a-z]+/) === null){
    return {value: '', error: 'Пароль должен содержать латинские буквы в нижнем регистре!'};
  }
  // if(value.match(/[!"№;%:\?\*()_\+`~@#\$\^&\-=]+/) === null){
  //   return {value: '', error: 'Пароль должен содержать хотя бы один спецсимвол!'};
  // }
  if(value.match(/^[a-zA-Z0-9]+$/) === null){ // !"№;%:\?\*()_\+`~@#\$\^&\-=
    return {value: '', error: 'Некорректный пароль!'};
  }
  if(value.length < 6){
    return {value: '', error: 'Слишком короткий пароль!'};
  }
  if(value.length > 30){
    return {value: '', error: 'Слишком длинный пароль!'};
  }
  return {value: value, error: ''};
}

export function ValidUserName(value: string): IValidAnswer {
  // Проверка на наличие только букв и пробелов
  if (value.match(/^[a-zA-Zа-яА-ЯёЁ\s]+$/) === null) {
      return { value: '', error: 'ФИО должно содержать только буквы и пробелы!' };
  }

  // Проверка на минимальную длину (например, 3 символа)
  if (value.length < 3) {
      return { value: '', error: 'ФИО слишком короткое!' };
  }

  // Проверка на максимальную длину (например, 100 символов)
  if (value.length > 100) {
      return { value: '', error: 'ФИО слишком длинное!' };
  }

  // Проверка на наличие хотя бы одного пробела (для разделения ФИО)
  if (!value.includes(' ')) {
      return { value: '', error: 'ФИО должно содержать хотя бы одно пробел для разделения!' };
  }

  return { value: value, error: '' };
}