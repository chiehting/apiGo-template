function redirect(path) {
    window.location = path;
}

// Example starter JavaScript for disabling form submissions if there are invalid fields
(function () {
'use strict'

  // Fetch all the forms we want to apply custom Bootstrap validation styles to
  var forms = document.querySelectorAll('.needs-validation')

  // Loop over them and prevent submission
  Array.prototype.slice.call(forms)
    .forEach(function (form) {
    form.addEventListener('submit', function (event) {
      if (!form.checkValidity()) {
      } else {
        var fn = window[form.id]()
        if (typeof fn === "function") fn();
        form.classList.add('was-validated')
      }
      event.preventDefault()
      event.stopPropagation()
    }, false)
    })
})()

// 用戶註冊
function registerForm() {
  let fromData = {
    "Username": document.getElementById("registerUsername").value,
    "Password": document.getElementById("registerPassword").value,
    "Email": document.getElementById("registerEmail").value,
    "FirstName": document.getElementById("registerFirstName").value,
    "LastName": document.getElementById("registerLastName").value,
    "Location": document.getElementById("registerLocation").value,
    "Organization": document.getElementById("registerOrganization").value,
    "PreferredLanguage": document.getElementById("registerLanguage").value,
  }

  fetch("./api/v1/user/register", {
    method: "POST",
    body: JSON.stringify(fromData)
  })
  .then(response => response.json())
  .then(function(data) {
    console.log(data)
    alert(JSON.stringify(data))
  })
  .catch(function(error) {
    console.log(error)
  });

  registerModal.toggle()
}

// 用戶登入
function signinForm() {
  let fromData = {
    "Account": document.getElementById("signinAccount").value,
    "Password": document.getElementById("signinPassword").value,
  }

  fetch("./api/v1/user/signin", {
    method: "POST",
    body: JSON.stringify(fromData)
  })
  .then(response => response.json())
  .then(function(data) {
    console.log(data)
    window.localStorage.setItem("token", data["result"]["token"])
    alert(JSON.stringify(data))
  })
  .catch(function(error) {
    console.log(error);
  });

  signinModal.toggle()
}

// 用戶資訊
function infoForm() {
  token = window.localStorage.getItem("token")
  if (token == "undefined") {
    alert("Token is undefined.")
    return
  }

  fetch("./api/v1/user", {
    method: "GET",
    headers: {
      'Content-Type': 'application/x-www-form-urlencoded',
      'Authorization': 'Bearer ' + token,
    },
  })
  .then(response => response.json())
  .then(function(data) {
    console.log(data)
    alert(JSON.stringify(data))
  })
  .catch(function(error) {
    console.log(error);
  });
  
  infoModal.toggle()
}

// 用戶登出
function signoutForm() {
  token = window.localStorage.getItem("token")
  if (token == "undefined") {
    alert("Token is undefined.")
    return
  }

  fetch("./api/v1/user/signout", {
    method: "POST",
    headers: {
      'Content-Type': 'application/x-www-form-urlencoded',
      'Authorization': 'Bearer ' + token,
    },
    body: {},
  })
  .then(response => response.json())
  .then(function(data) {
    console.log(data)
    window.localStorage.removeItem("token")
    alert(JSON.stringify(data))
  })
  .catch(function(error) {
    console.log(error);
  });

  signoutModal.toggle()
}

var registerModal = new bootstrap.Modal(document.getElementById('registerModal'), {})
var signinModal = new bootstrap.Modal(document.getElementById('signinModal'), {})
var infoModal = new bootstrap.Modal(document.getElementById('infoModal'), {})
var signoutModal = new bootstrap.Modal(document.getElementById('signoutModal'), {})
