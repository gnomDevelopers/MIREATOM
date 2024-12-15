import axios from "axios";
import {
  API, DEVMODE, GET_COOKIE,
  type IAPI_Login,
  type IAPI_Register,
  type Article,
} from "@/helpers/constants";

//////// ПОЛЬЗОВАТЕЛЬ ////////

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

//получение информации о пользователе
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

//////// СТАТЬИ ////////

//получение всех статей
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

//получение файла статьи
export function API_Article_Get_File(articleId: number){
  return new Promise((resolve, reject) => {
    axios.get(`${API}/article/file/${articleId}`)
    .then(response => {
      if(DEVMODE) console.log('Articles file get success: ', response);
      resolve(response);
    })
    .catch(error => {
      if(DEVMODE) console.log('Articles file get error: ', error);
      reject(error);
    })
  });
};

//получение статьи пользователя
export function API_Article_Get_ByID(userID: number): Promise<Article[]>{
  return new Promise((resolve, reject) => {
    axios.get(`${API}/article/user_id/${userID}`)
    .then(response => {
      if(DEVMODE) console.log('Article get by id success: ', response);
      resolve(response.data);
    })
    .catch(error => {
      if(DEVMODE) console.log('Article get by id error: ', error);
      reject(error);
    })
  });
};

//добавление статьи пользователем
export function API_Article_Post(formData: FormData): Promise<Article[]>{
  return new Promise((resolve, reject) => {
    axios.post(`${API}/auth/article`, formData, {
      headers: {
        Authorization: 'Bearer ' + GET_COOKIE('access_token'),
      }
    })
    .then(response => {
      if(DEVMODE) console.log('Article create success: ', response);
      resolve(response.data);
    })
    .catch(error => {
      if(DEVMODE) console.log('Article create error: ', error);
      reject(error);
    });
  });
};

//////// ФОРМУЛЫ ////////

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

//обновление формулы
export function API_Update_Formula(data: {id: number, title: string, value: string}){
  return new Promise((resolve, reject) => {
    axios.put(`${API}/auth/formula`, data, {
      headers: {
        Authorization: 'Bearer ' + GET_COOKIE('access_token'),
      }
    })
    .then(response => {
      if(DEVMODE) console.log('Update formula success: ', response);
      resolve(response);
    })
    .catch(error => {
      if(DEVMODE) console.log('Update formula error: ', error);
      reject(error);
    })
  });
};

//получение истории коммитов формулы
export function API_Get_Formula_Commits(formulaID: number){
  return new Promise((resolve, reject) => {
    axios.get(`${API}/auth/formula/id/${formulaID}/commits`, {
      headers: {
        Authorization: 'Bearer ' + GET_COOKIE('access_token'),
      }
    })
    .then(response => {
      if(DEVMODE) console.log('Get formula commits success: ', response);
      resolve(response);
    })
    .catch(error => {
      if(DEVMODE) console.log('Get formula commits error: ', error);
      reject(error);
    })
  });
};

//распознавание формулы с фото
export function API_Get_Formula_From_Photo(formData: FormData){
  return new Promise((resolve, reject) => {
    axios.post(`${API}/formula/recognize/`, formData)
    .then(response => {
      if(DEVMODE) console.log('Get formula from photo success: ', response);
      resolve(response);
    })
    .catch(error => {
      if(DEVMODE) console.log('Get formula from photo error: ', error);
      reject(error);
    })
  });
};

//получение уникальности формулы
export function API_Get_Formula_Analysis(formData: FormData){
  return new Promise((resolve, reject) => {
    axios.post(`${API}/formula/analysis/`, formData)
    .then(response => {
      if(DEVMODE) console.log('Get formula analysis success: ', response);
      resolve(response);
    })
    .catch(error => {
      if(DEVMODE) console.log('Get formula analysis error: ', error);
      reject(error);
    })
  });
};