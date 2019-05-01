"use strict";

// DOMCONTENTLOADED to display basic infos
document.addEventListener("DOMContentLoaded", displayAllData);

function displayAllData() {
  // Get kasten counter
  let indexKastenCounterPublic = document.getElementById(
    "indexKastenCounterPublic"
  );
  // Prepare http request
  let xhr3 = new XMLHttpRequest();
  xhr3.responseType = `json`;
  let url = "http://localhost:8080/karteikasten/";
  xhr3.open("GET", url, true);
  xhr3.setRequestHeader("Content-Type", "application/json");
  // Callback function
  xhr3.onreadystatechange = function() {
    if (xhr3.readyState == 4 && xhr3.status == 200) {
      // Count all public Karteikaesten
      let counterPublic = 0;
      for (let i = 0; i < xhr3.response.length; i++) {
        if (xhr3.response[i].private == "false") {
          counterPublic++;
        }
      }
      indexKastenCounterPublic.innerHTML = counterPublic;
    }
  };
  // GET data
  xhr3.send(null);
}

// REGISTER
let registerButton = document.getElementById("registerButton");
registerButton.addEventListener("click", registerNewUser);

function registerNewUser(e) {
  // Clear all register alerts
  document.getElementById("registerUsernameError").innerHTML = "";
  document
    .getElementById("registerUsernameError")
    .classList.remove("registerAlert");
  document.getElementById("registerEmailError").innerHTML = "";
  document
    .getElementById("registerEmailError")
    .classList.remove("registerAlert");
  document.getElementById("registerPasswoerterError").innerHTML = "";
  document
    .getElementById("registerPasswoerterError")
    .classList.remove("registerAlert");
  document.getElementById("registerCheckboxError").innerHTML = "";
  document
    .getElementById("registerCheckboxError")
    .classList.remove("registerAlert");
  e.preventDefault();
  // Get values from form
  let username = document.getElementById("registerUsername").value;
  let email = document.getElementById("registerEmail").value;
  let passwordOne = document.getElementById("registerPasswordOne").value;
  let passwordTwo = document.getElementById("registerPasswordTwo").value;
  // Check if both passwords are the same
  if (passwordOne !== passwordTwo) {
    let registerFailPassword = document.getElementById(
      "registerPasswoerterError"
    );
    registerFailPassword.innerHTML = "Die Passwörter stimmen nicht überein!";
    registerFailPassword.classList.add("registerAlert");
    // Check if checkbox is checked
  } else if (!document.getElementById("datenschutzCheckbox").checked) {
    let checkboxNotChecked = document.getElementById("registerCheckboxError");
    checkboxNotChecked.innerHTML = "Bitte setzen sie hier einen Haken!";
    checkboxNotChecked.classList.add("registerAlert");
  } else {
    // Prepare http request
    let xhr = new XMLHttpRequest();
    let url = "http://localhost:8080/user/";
    xhr.open("POST", url, true);
    xhr.setRequestHeader("Content-Type", "application/json");
    // Callback function
    xhr.onreadystatechange = function() {
      if (xhr.readyState == 4 && xhr.status == 200) {
        // Check if username exists
        if (JSON.parse(xhr.response) == "Username already exist") {
          let usernameExists = document.getElementById("registerUsernameError");
          usernameExists.innerHTML = "Dieser Nutzername ist bereits vergeben!";
          usernameExists.classList.add("registerAlert");
          // Check if email exists
        } else if (JSON.parse(xhr.response) == "Email already exist") {
          let emailExists = document.getElementById("registerEmailError");
          emailExists.innerHTML = "Diese Mail ist bereits vergeben!";
          emailExists.classList.add("registerAlert");
        } else {
          // Register new user and login
          let xhr4 = new XMLHttpRequest();
          url = "http://localhost:8080/login/";
          xhr4.open("POST", url, true);
          xhr4.setRequestHeader("Content-Type", "application/json");
          // Callback function
          xhr4.onreadystatechange = function() {
            if (xhr4.readyState == 4 && xhr4.status == 200) {
              if (xhr4.response) {
                window.location.href =
                  "http://localhost:8080/static/meinekarteien.html";
              }
            }
          };
          // POST data
          let data = JSON.stringify({
            username: username,
            password: passwordOne
          });
          xhr4.send(data);
        }
      }
    };
    // POST data
    let data = JSON.stringify({
      username: username,
      email: email,
      password: passwordOne
    });
    xhr.send(data);
  }
}

// LOGIN
document.getElementById("loginButton").addEventListener("click", login);

function login(e) {
  e.preventDefault();
  // Get values from form
  let username = document.getElementById("loginUsername").value;
  let password = document.getElementById("loginPassword").value;
  // Check if both inputs are filled
  if (username === "" || password === "") {
    alert("Bitte beide Felder ausfüllen.");
  } else {
    // Prepare http request
    let xhr4 = new XMLHttpRequest();
    let url = "http://localhost:8080/login/";
    xhr4.open("POST", url, true);
    xhr4.setRequestHeader("Content-Type", "application/json");
    // Callback function
    xhr4.onreadystatechange = function() {
      if (xhr4.readyState == 4 && xhr4.status == 200) {
        if (xhr4.response) {
          window.location.href =
            "http://localhost:8080/static/meinekarteien.html";
        } else {
          alert("Nutzername oder Passwort sind falsch!");
        }
      }
    };
    // POST data
    let data = JSON.stringify({
      username: username,
      password: password
    });
    xhr4.send(data);
    // Clear forms after submit
    document.getElementById("loginUsername").value = "";
    document.getElementById("loginPassword").value = "";
  }
}
