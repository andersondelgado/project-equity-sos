import axios from 'axios';


axios.defaults.headers.common['X-Requested-With'] = 'XMLHttpRequest';
const token = localStorage.getItem('token');
if (token) {
  axios.defaults.headers.common.Authorization = 'Bearer ' + token;
}


