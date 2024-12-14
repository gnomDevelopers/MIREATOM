import axios from "axios";
import {
  API, DEVMODE, GET_COOKIE,
  type IAPI_Login,
  type IAPI_Register,
  type Article,
} from "@/helpers/constants";

//проверка аутентификации пользователя
export function API_Authenticate(){
  return new Promise((resolve, reject) => {
    axios.get(`${API}/login`,  {
      headers: {
        Authorization: 'Bearer ' + GET_COOKIE('access_token'),
      }
     })
    .then(response => {
      if(DEVMODE) console.log('Authentication success: ', response);
      resolve(response);
    })
    .catch(error => {
      if(DEVMODE) console.log('Authentication error: ', error);
      reject(error);
    });
  });
};

//вход в аккаунт
export function API_Login(data: IAPI_Login){
  return new Promise((resolve, reject) => {
    axios.post(`${API}/login`, data)
    .then(response => {
      if(DEVMODE) console.log('Login post success: ', response);
      resolve(response);
    })
    .catch(error => {
      if(DEVMODE) console.log('Login post error: ', error);
      reject(error);
    });
  });
};

//выход из аккаунта
export function API_Logout(){
  return new Promise((resolve, reject) => {
    axios.post(`${API}/logout`)
    .then(response => {
      if(DEVMODE) console.log('Logout post success: ', response);
      resolve(response);
    })
    .catch(error => {
      if(DEVMODE) console.log('Logout post error: ', error);
      reject(error);
    });
  });
};

//тест запрос
export function API_Health(){
  return new Promise((resolve, reject) => {
    axios.get(`${API}/health`)
    .then(response => {
      if(DEVMODE) console.log('Health get success: ', response);
      resolve(response);
    })
    .catch(error => {
      if(DEVMODE) console.log('Health get error: ', error);
      reject(error);
    });
  });
};

//регистрация
export function API_Register(data: IAPI_Register){
  return new Promise((resolve, reject) => {
    axios.post(`${API}/signup`, data)
    .then(response => {
      if(DEVMODE) console.log('Sign up post success: ', response);
      resolve(response);
    })
    .catch(error => {
      if(DEVMODE) console.log('Sign up post error: ', error);
      reject(error);
    });
  });
};

//вход в аккаунт
export function API_Get_User_Data(userID: number){
  return new Promise((resolve, reject) => {
    axios.get(`${API}/user/${userID}`)
    .then(response => {
      if(DEVMODE) console.log('User data get success: ', response);
      resolve(response);
    })
    .catch(error => {
      if(DEVMODE) console.log('User data get error: ', error);
      reject(error);
    });
  });
};

//сохранение формулы
export function API_Save_Formula(data:{title: string, value: string}){
  return new Promise((resolve, reject) => {
    axios.post(`${API}/auth/formula`, data, {
      headers: {
        Authorization: 'Bearer ' + GET_COOKIE('access_token'),
      }
    })
    .then(response => {
      if(DEVMODE) console.log('Save formula success: ', response);
      resolve(response);
    })
    .catch(error => {
      if(DEVMODE) console.log('Save formula error: ', error);
      reject(error);
    });
  });
};

//получение истории формул
export function API_Get_Formuls_History(userID: number, page: number){
  return new Promise((resolve, reject) => {
    axios.get(`${API}/formula/history/user/${userID}/page/${page}`)
    .then(response => {
      if(DEVMODE) console.log('Get formula history success: ', response);
      resolve(response);
    })
    .catch(error => {
      if(DEVMODE) console.log('Get formula history error: ', error);
      reject(error);
    });
  });
};

//////// СТАТЬИ ////////
export function API_Articles_Get(): Promise<Article[]>{
  return new Promise((resolve, reject) => {
    axios.get(`${API}/article`)
    .then(response => {
      if(DEVMODE) console.log('Articles get success: ', response);
      resolve(response.data as Article[]);
    })
    .catch(error => {
      if(DEVMODE) console.log('Articles get error: ', error);
      reject(error);
    })
  });
};

export function API_ArticleFile_Get(articleId: number){
  return new Promise((resolve, reject) => {
    axios.get(`${API}/auth/article/file/${articleId}`)
    .then(response => {
      if(DEVMODE) console.log('Articles file get success: ', response);
      resolve(response);
    })
    .catch(error => {
      if(DEVMODE) console.log('Articles file error: ', error);
      reject(error);
    })
  });
};