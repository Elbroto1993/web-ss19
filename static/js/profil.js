"use strict";

document.addEventListener("DOMContentLoaded", displayAllData);

// Get data for karteikasten sidebar counter
function displayAllData() {
  // Get kasten counter
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
      document.getElementById(
        "indexKastenCounterPublic"
      ).innerHTML = counterPublic;
    }
  };
  // GET data
  xhr3.send(null);
  // Get user data
  let xhr6 = new XMLHttpRequest();
  url = "http://localhost:8080/user/id";
  xhr6.open("GET", url, true);
  xhr6.setRequestHeader("Content-Type", "application/json");
  // Callback function
  xhr6.onreadystatechange = function() {
    if (xhr6.readyState == 4 && xhr6.status == 200) {
      let res = JSON.parse(xhr6.response);
      document.getElementById("logoutUsername").innerHTML = res.username;
      document.getElementById("profilUsernameOutput").innerHTML = res.username;
      document.getElementById("profilEmailOutput").innerHTML = res.email;
      document.getElementById("profilDatumOutput").innerHTML = res.createdat;
    }
  };
  // GET data
  xhr6.send(null);
  // Get created kaesten from user
  let xhr7 = new XMLHttpRequest();
  url = "http://localhost:8080/karteikasten/createdbyuser";
  xhr7.open("GET", url, true);
  xhr7.setRequestHeader("Content-Type", "application/json");
  // Callback function
  xhr7.onreadystatechange = function() {
    if (xhr7.readyState == 4 && xhr7.status == 200) {
      let res = JSON.parse(xhr7.responseText);
      document.getElementById("profilKastenOutput").innerHTML = res.length;
      document.getElementById("indexKartenCounterUser").innerHTML = res.length;
    }
  };
  // GET data
  xhr7.send(null);
}
// LOGOUT
document.getElementById("logoutButton").addEventListener("click", logout);

function logout() {
  let xhr = new XMLHttpRequest();
  let url = "http://localhost:8080/login/logout";
  xhr.open("GET", url, true);
  // Callback function
  xhr.onreadystatechange = function() {
    if (xhr.readyState == 4 && xhr.status == 200) {
      window.location.href = "http://localhost:8080/static/";
    }
  };
  // GET data
  xhr.send(null);
}

// UPDATE PROFILE
document.getElementById("updateProfilButton").addEventListener("click", update);

function update(e) {
  e.preventDefault();
  // Get all necessary elements
  let email = document.getElementById("updateProfilEmail").value;
  let oldPassword = document.getElementById("updateProfilAltesPasswort").value;
  let passwordOne = document.getElementById("updateProfilNeuesPasswortOne")
    .value;
  let passwordTwo = document.getElementById("updateProfilNeuesPasswortTwo")
    .value;
  let passwordError = document.getElementById("neuesPasswort");
  let emailError = document.getElementById("neueEmail");
  // Clear all alerts
  passwordError.innerHTML = "";
  passwordError.classList.remove("registerAlert");
  emailError.innerHTML = "";
  emailError.classList.remove("registerAlert");
  // Check if password inputs are correct
  if (passwordOne !== passwordTwo) {
    passwordError.innerHTML = "Die Passwörter stimmen nicht überein!";
    passwordError.classList.add("registerAlert");
  } else if (passwordOne === oldPassword) {
    passwordError.innerHTML =
      "Altes Passwort darf nicht mit neuem übereinstimmen!";
    passwordError.classList.add("registerAlert");
  } else {
    // Prüfen ob altes passwort mit der datenbank übereinstimmt, sonst fehlermeldung
    let xhr2 = new XMLHttpRequest();
    let url = `http://localhost:8080/user/id`;
    xhr2.responseType = `json`;
    xhr2.open("GET", url, true);
    xhr2.onreadystatechange = function() {
      if (xhr2.readyState == 4 && xhr2.status == 200) {
        if (xhr2.response.password == oldPassword) {
          let xhr = new XMLHttpRequest();
          let url = "http://localhost:8080/user/update";
          xhr.open("POST", url, true);
          xhr.setRequestHeader("Content-Type", "application/json");
          xhr.onreadystatechange = function() {
            if (xhr.readyState == 4 && xhr.status == 200) {
              if (JSON.parse(xhr.response) == "Email already exist") {
                emailError.innerHTML = "Diese Mail ist bereits vergeben!";
                emailError.classList.add("registerAlert");
              } else {
                console.log(xhr.responseText);
                // Clear forms
                document.getElementById("updateProfilEmail").value = "";
                document.getElementById("updateProfilAltesPasswort").value = "";
                document.getElementById("updateProfilNeuesPasswortOne").value =
                  "";
                document.getElementById("updateProfilNeuesPasswortTwo").value =
                  "";
              }
            }
          };
          let data = JSON.stringify({
            email: email,
            password: passwordOne
          });
          xhr.send(data);
        } else {
          passwordError.innerHTML = "Altes Passwort ist falsch!";
          passwordError.classList.add("registerAlert");
        }
      }
    };
    xhr2.send(null);
  }
}

// DELETE PROFILE
// Toggle modal for second delete button
document
  .getElementById("profilDeleteButton")
  .addEventListener("click", toggleDeleteModal);

function toggleDeleteModal(e) {
  e.preventDefault(e);
  document.getElementById("profilDeleteModal").classList.add("is-active");
}

document
  .getElementById("modal-close")
  .addEventListener("click", untoggleDeleteModal);
document
  .getElementById("profilKeepButton")
  .addEventListener("click", untoggleDeleteModal);

function untoggleDeleteModal() {
  document.getElementById("profilDeleteModal").classList.remove("is-active");
}
// Final delete button
document
  .getElementById("profilFinallyDeleteButton")
  .addEventListener("click", deleteProfile);
function deleteProfile() {
  // Delete profile
  let xhr = new XMLHttpRequest();
  let url = "http://localhost:8080/user/";
  xhr.open("DELETE", url, true);
  xhr.onreadystatechange = function() {
    if (xhr.readyState == 4 && xhr.status == 200) {
      window.location.href = "http://localhost:8080/static/";
      console.log(xhr.responseText);
    }
  };
  xhr.send(null);
  // Then logout for security reasons
  let xhr2 = new XMLHttpRequest();
  let url2 = "http://localhost:8080/login/logout";
  xhr2.open("GET", url2, true);
  // Callback function
  xhr2.onreadystatechange = function() {
    if (xhr2.readyState == 4 && xhr2.status == 200) {
      console.log(xhr2.responseText);
    }
  };
  // GET data
  xhr2.send(null);
}

// Add eventListener to new kartei button
document.getElementById("newKarteiButton").addEventListener("click", newKartei);
// NEW KARTEI
function newKartei() {
  location.href = "edit.html";
}

// IMAGE UPLOAD
let imageInput = document.getElementById("imageUploadInput");
let imageButton = document.getElementById("imageUploadButton");
imageButton.addEventListener("click", chooseImage);
imageInput.addEventListener("change", uploadImage);

function chooseImage(e) {
  e.preventDefault();
  imageInput.click();
}
function uploadImage(e) {
  let a = e.target.files[0];
  console.log(a);
}
