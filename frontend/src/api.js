const API_ENDPOINT = process.env.BACKEND_API_BASE;

export const getPrivateMessage = function(idToken) {
  return fetch(`${API_ENDPOINT}/private`, {
    method: "get",
    headers: new Headers({
      Authorization: `Bearer ${idToken}`
    }),
    credentials: "same-origin"
  }).then(res => {
    if (res.ok) {
      return res.json();
    } else {
      throw Error(`Request rejected with status ${res.status}`);
    }
  });
};

export const getPublicMessage = function() {
  return fetch(`${API_ENDPOINT}/public`);
};


export const getArticles = function(id) {
  return fetch(`${API_ENDPOINT}/articles/${id}`, {
    method: "get",
    // headers: new Headers({
    //   Authorization: `Bearer ${idToken}`
    // }),
    // credentials: "same-origin"
  }).then(res => {
    if (res.ok) {
      return res.json();
    } else {
      throw Error(`Request rejected with status ${res.status}`);
    }
  });
};

export const createArticle = function(data, idToken) {
  return fetch(`${API_ENDPOINT}/articles`, {
    method: "POST",
    headers: new Headers({
      "Content-Type": "application/json; charset=utf-8",
      Authorization: `Bearer ${idToken}`,
      mode: 'cors',
      "Access-Control-Allow-Origin": "*",
    }),
    body: JSON.stringify(data),
    credentials: "same-origin"
  }).then(res => {
    if (res.ok) {
      return res.json();
    } else {
      throw Error(`Request rejected with status ${res.status}`);
    }
  });
};