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
  if(value.match(/^[a-z_]+@[a-z\.]+\.[a-z]{2,6}$/) === null) {
    return {value: '', error: 'Некорректный логин!'};
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
  const fio = value.trim().split(' ');
  if(fio.length !== 3){
    return {value: '', error: 'Некорректный фио пользователя!'};
  }
  // Проверка на корректные символы: буквы и цифры
  if (value.match(/^[a-zA-Zа-яА-ЯёЁ\s]+$/) === null) {
    return { value: '', error: 'ФИО может содержать только буквы!' };
  }
  return {value: value, error: ''};
}