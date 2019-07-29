import { Subject } from 'rxjs';

const subject = new Subject();

export const messageService = {
    setSearchs: (message: any) => subject.next({ text: message }),
    clearSearchs: () => subject.next(),
    getSearchs: () => subject.asObservable()
};

export const TableService = {
  setTable: (tableData:any)=> subject.next({ array: tableData }),
  clearTable: () => subject.next(),
  getTable: () => subject.asObservable()
};


export default class Global {

  public static DOMAIN = 'http://localhost:8000/api/';
  // public static DOMAIN = 'https://equity-sos-go-dev-grouchy-swan.mybluemix.net/api/';

  public static const = {
    LOGIN: Global.DOMAIN + 'login',
    REGISTER: Global.DOMAIN + 'register',     
    COUNTRY: Global.DOMAIN + 'country/all',
    CHANGE_PASSWORD: Global.DOMAIN + 'reset-password',
    FORGOT_PASSWORD: Global.DOMAIN + 'forgot-password',
    RESET_TOKEN_PASSWORD: Global.DOMAIN + 'reset-token-password',
    MENU: Global.DOMAIN + 'menu',
    USER_INFO: Global.DOMAIN + 'user-info',
    PRUEBA_ALL: Global.DOMAIN + 'test/all',
    PRUEBA_PAGINATE: Global.DOMAIN + 'test/paginate/',
    PRUEBA_SEARCH_PAGINATE: Global.DOMAIN + 'test/search-paginate/',
    PRUEBA_EDIT: Global.DOMAIN + 'test/edit/',
    PRUEBA_SAVE: Global.DOMAIN + 'test/add',
    PRUEBA_UPDATE: Global.DOMAIN + 'test/update/',
    PRUEBA_DELETE: Global.DOMAIN + 'test/delete/',
    ARTICLE_ALL: Global.DOMAIN + 'article/all',
    PERMISSION_LIST: Global.DOMAIN + 'permission/list',
    ARTICLE_PAGINATE: Global.DOMAIN + 'article/paginate/',
    ARTICLE_SEARCH_PAGINATE: Global.DOMAIN + 'article/search-paginate/',
    ARTICLE_EDIT: Global.DOMAIN + 'article/edit/',
    ARTICLE_SAVE: Global.DOMAIN + 'article/add',
    ARTICLE_UPDATE: Global.DOMAIN + 'article/update/',
    ARTICLE_DELETE: Global.DOMAIN + 'article/delete/',
    USERS_SAVE: Global.DOMAIN + 'users/add',
    USERS_EDIT: Global.DOMAIN + 'users/edit/',
    USERS_UPDATE: Global.DOMAIN + 'users/update/',
    USERS_DELETE: Global.DOMAIN + 'users/delete/',
    USERS_PAGINATE: Global.DOMAIN + 'users/paginate/',
    USERS_SEARCH_PAGINATE: Global.DOMAIN + 'users/search-paginate/',
    CATEGORY_ALL: Global.DOMAIN + 'category/all',
    CATEGORY_PAGINATE: Global.DOMAIN + 'category/paginate/',
    CATEGORY_SEARCH_PAGINATE: Global.DOMAIN + 'category/search-paginate/',
    CATEGORY_EDIT: Global.DOMAIN + 'category/edit/',
    CATEGORY_SAVE: Global.DOMAIN + 'category/add',
    CATEGORY_UPDATE: Global.DOMAIN + 'category/update/',
    CATEGORY_DELETE: Global.DOMAIN + 'category/delete/',
    POST_SAVE: Global.DOMAIN + 'post/add',
    POST_PAGINATE: Global.DOMAIN + 'post/paginate/',
    POST_SEARCH_PAGINATE: Global.DOMAIN + 'post/search-paginate/',
    POST_BUSINESS_PAGINATE: Global.DOMAIN + 'post/business/paginate/',
    POST_BUSINESS_SEARCH_PAGINATE: Global.DOMAIN + 'post/business/search-paginate/',
    POST_EDIT: Global.DOMAIN + 'post/edit/',
    POST_UPDATE: Global.DOMAIN + 'post/update/',
    POST_BUSINESS_EDIT: Global.DOMAIN + 'post/business/edit/',
    POST_BUSINESS_UPDATE: Global.DOMAIN + 'post/business/update/',
    GET_COUNTRY_LIST: Global.DOMAIN + 'country/all',
    // Kyc
    GET_KYC_LIST: Global.DOMAIN + 'kyc/all',
    GET_KYC_USER: Global.DOMAIN + 'kyc/userKyc',
    GET_KYC_USER_LIST: Global.DOMAIN + 'kyc-user/{userId}',
    POST_KYC_USER_DOCUMENTS: Global.DOMAIN + 'kyc-user/add',
    PUT_KYC_USER_DOCUMENTS: Global.DOMAIN + 'kyc-user/edit',
    PUT_KYC_USER_DOCUMENTS_VALIDATE: Global.DOMAIN + 'kyc/adminValidate'
  };

  public static headers: any = {
    "Content-Type": "application/json"
  };
  
  public static cors: any = {
    headers: Global.headers
  }

  public getBase64ImageEncode(img: any): any {
    const promise: any = new Promise((resolve, reject) => {
      let baseString: any;
      const reader: any = new FileReader();
      reader.readAsDataURL(img);
      reader.onload = (() => {
        baseString = reader.result;
        return resolve(baseString);
      });
      reader.onerror = ((error: any) => {
        // console.log('Error: ', error);
        // return reject()
      });
    });
    return promise;
  }

  public crud() {
    let crud = ['read', 'create', 'edit', 'update', 'delete'];
    return crud;
  }

  public perm(url: any) {
    let crud = [];
    const perm: any = localStorage.getItem('permission');
    const data = JSON.parse(perm);
    if (data !== null || data !== undefined) {
      for (var i = 0; i < data.length; i++) {
        if (data[i].url === url) {
          crud = data[i].value;
        }
      }
    }
    return crud;
  }

  public multiFileBase64(event: any) {
    let file = event.target.files;
    let arr: any = [];

    for (let i = 0; i < file.length; i++) {
      let f = file[i];
      let reader = new FileReader();
      reader.readAsDataURL(f);
      let baseString;
      reader.onloadend = ((e: any) => {
        baseString = e.target.result;
        arr.push(baseString);
      });
    }
    return arr;
  }

  public objectFindByKey(array:any, key:any, value:any) {
    for (let i = 0; i < array.length; i++) {
      if (array[i][key] === value) {
        return array[i];
      }
    }
    return null;
  }

  public isRead(url: any) {
    let per = this.perm(url);
    let index = per.indexOf(this.crud()[0])
    let crud;
    crud = (per[index] !== undefined) ? per[index] : 0;
    return crud;
  }

  public isCreate(url: any) {
    let per = this.perm(url);
    let index = per.indexOf(this.crud()[1])
    let crud;
    crud = (per[index] !== undefined) ? per[index] : 0;
    return crud;
  }

  public isEdit(url: any) {
    let per = this.perm(url);
    let index = per.indexOf(this.crud()[2])
    let crud;
    crud = (per[index] !== undefined) ? per[index] : 0;
    return crud;
  }

  public isUpdate(url: any) {
    let per = this.perm(url);
    let index = per.indexOf(this.crud()[3])
    let crud;
    crud = (per[index] !== undefined) ? per[index] : 0;
    return crud;
  }

  public isDelete(url: any) {
    let per = this.perm(url);
    let index = per.indexOf(this.crud()[4])
    let crud;
    crud = (per[index] !== undefined) ? per[index] : 0;
    return crud;
  }

}
