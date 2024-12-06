//all constants, interfaces, types, etc.

//contants

//api
export const API = '/api';
export const DEVMODE = true;
//status window
export const STATUS_WINDOW_TIME = 3000;

//statusWindow codes
export enum StatusCodes {
  'error', 'info', 'loading', 'success'
};

//interfaces

//statusWindow interface
export interface IStatusWindow{
  id: number,
  status: StatusCodes,
  text: string,
  time: number,
};

export interface IAPI_Login{
  login: string,
  password: string,
};

//types

export type TMaybeNumber = number | null;
export type TMaybeBoolean = boolean | null;
export type TMaybeString = string | null;

// functions

export function CHECK_COOKIE_EXIST(cookieName: string): boolean{
  const allCookie = document.cookie.split(';');
  for(let cookie of allCookie){
    if(cookie.split('=')[0].trim() === cookieName) return true;
  }
  return false;
}

export function GET_COOKIE(cookieName: string): string{
  const allCookie = document.cookie.split(';');
  for(let cookie of allCookie){
    let [key, value] = cookie.split('=');
    if(key.trim() === cookieName) return value;
  }
  return '';
}