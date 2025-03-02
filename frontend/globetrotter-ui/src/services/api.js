import axios from "axios";

export const postRequestToJsonAPI = (url, payload) => (new Promise((resolve, reject) => {
    const headers = {
      'Content-Type': 'application/json',
      accept: 'application/json',
    }
    // const token = window.sessionStorage.getItem('token');
    // headers.Authorization = `Bearer ${token}`;
    console.log("validation payload api", payload);
    axios.post(url, payload, {headers: headers, maxContentLength: 9000000000,timeout:15000})
      .then((response) => {
        if (response.status === 200) {
          if (response.data) {
            resolve({ status: response.status, data: response.data, obj_len: response.data.length });
          } else {
            resolve({ status: response.status });
          }
        } else if (response.data) {
          reject({ status: response.status, data: response.data });
        } else {
          reject({ status: response.status });
        }
      }).catch((error) => {
        if (error.response && error.response.data) {
          reject({ status: 422, data: error.response.data });
        } else {
          reject({ status: 422 });
        }
      });
  }));

  export const getRequestToAPI = (url) => (new Promise((resolve, reject) => {
    const headers = {
      'Content-Type': 'application/json',
    }
    //const token = window.sessionStorage.getItem('token');
  // const token = `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCIsImlzcyI6IjlWazFwOE5xOGJzQ0tNTUNtUVY1dU9LeXNuWllRd1ZiIn0.eyJ1c2VybmFtZSI6ImxpdmUtZGViLTIwMCIsInJvbGUiOiJ1c2VyIiwibmFtZSI6IlhZWiIsImlhdCI6MTcwMDM5MzM2NX0.eYdwpVXWuFezee6ZUppSaJjuNR2amSNkzGt23jNb2_s`
  
    //headers.Authorization = `Bearer ${token}`;
    axios.get(url, {headers: headers})
      .then((response) => {
        if (response.status === 200) {
          if (response.data) {
            resolve({ status: response.status, data: response.data, obj_len: response.data.length });
          } else {
            resolve({ status: response.status });
          }
        } else if (response.data) {
          reject({ status: response.status, data: response.data });
        } else {
          reject({ status: response.status });
        }
      }).catch((error) => {
        if (error.response && error.response.data) {
          reject({ status: 422, data: error.response.data });
        } else {
          reject({ status: 422 });
        }
      });
  }));