"use strict";

document.addEventListener("DOMContentLoaded", displayAllData);

function displayAllData() {
  // CHECK IF USER IS LOGGED IN
  let xhr5 = new XMLHttpRequest();
  let url2 = "http://localhost:8080/login/checkLoggedIn";
  xhr5.responseType = `json`;
  xhr5.open("GET", url2, true);
  xhr5.setRequestHeader("Content-Type", "application/json");
  // Callback function
  xhr5.onreadystatechange = function() {
    if (xhr5.readyState == 4 && xhr5.status == 200) {
      // If user is logged in add sidebar items and change login field
      if (xhr5.response) {
        // Add sidebar items
        document.getElementById("menuList").innerHTML += `<li class="menuItems">
        <a href="meinekarteien.html"
          ><img
            src="icons/Meine-Karteien.svg"
            alt=""
            height="30"
            width="30"
          />
          Meine Karteien<small class="karteiZaehler" id="indexKartenCounterUser">0</small></a
        >
        </li>
        <li class="menuItems">
          <a href="profil.html"
            ><img
              src="icons/Mein-Profil.svg"
              alt=""
              height="30"
              width="30"
            />
            Mein Profil</a
          >
        </li>`;
        // Change login field
        document.getElementById("loginField").innerHTML = `<div class="logout">
        <button class="btnYellow newKarteiButton" id="newKarteiButton">
          Neue Kartei
        </button>
        <div class="headerUserInfo">
          <h4 id="logoutUsername">Max Mustermann</h4>
          <button class="logoutButton" id="logoutButton">
           <small>Logout</small>
          </button>
        </div>
          <img
          src="images/defaultUser.png"
          alt="gray default male profile picture"
          height="30"
          width="30"
          />
        </div>`;
        // Change name in logout field
        // Prepare http request
        let xhr6 = new XMLHttpRequest();
        url = "http://localhost:8080/user/id";
        xhr6.open("GET", url, true);
        //xhr6.responseType = `json`;
        xhr6.setRequestHeader("Content-Type", "application/json");
        // Callback function
        xhr6.onreadystatechange = function() {
          if (xhr6.readyState == 4 && xhr6.status == 200) {
            let res = JSON.parse(xhr6.response);
            document.getElementById("logoutUsername").innerHTML = res.username;
          }
        };
        // GET data
        xhr6.send(null);
        // Add eventListener to logout button
        document
          .getElementById("logoutButton")
          .addEventListener("click", logout);
        // Add eventListener to new kartei button
        document
          .getElementById("newKarteiButton")
          .addEventListener("click", newKartei);
        // Get user karten counter
        let indexKartenCounterUser = document.getElementById(
          "indexKartenCounterUser"
        );
        // Prepare http request
        let xhr4 = new XMLHttpRequest();
        xhr4.responseType = `json`;
        url = "http://localhost:8080/karteikasten/user";
        xhr4.open("GET", url, true);
        xhr4.setRequestHeader("Content-Type", "application/json");
        // Callback function
        xhr4.onreadystatechange = function() {
          if (xhr4.readyState == 4 && xhr4.status == 200) {
            if (xhr4.response.length != null) {
              indexKartenCounterUser.innerHTML = xhr4.response.length;
            }
          }
        };
        // GET data
        xhr4.send(null);
      }
    }
  };
  // THE FOLLOWING ALSO HAPPENS FOR USERS WHO AREN'T LOGGED IN
  // GET data
  xhr5.send(null);
  ////
  // Get users counter
  let indexNutzerCounter = document.getElementById("indexNutzerCounter");
  // Prepare http request
  let xhr = new XMLHttpRequest();
  xhr.responseType = `json`;
  let url = "http://localhost:8080/user/getAllUsersLength";
  xhr.open("GET", url, true);
  xhr.setRequestHeader("Content-Type", "application/json");
  // Callback function
  xhr.onreadystatechange = function() {
    if (xhr.readyState == 4 && xhr.status == 200) {
      indexNutzerCounter.innerHTML = xhr.response;
    }
  };
  // GET data
  xhr.send(null);

  // Get all karten counter
  let indexKartenCounter = document.getElementById("indexKartenCounter");
  // Prepare http request
  let xhr2 = new XMLHttpRequest();
  xhr2.responseType = `json`;
  url = "http://localhost:8080/karteikarte/";
  xhr2.open("GET", url, true);
  xhr2.setRequestHeader("Content-Type", "application/json");
  // Callback function
  xhr2.onreadystatechange = function() {
    if (xhr2.readyState == 4 && xhr2.status == 200) {
      indexKartenCounter.innerHTML = xhr2.response.length;
    }
  };
  // GET data
  xhr2.send(null);

  // Get kasten counter
  let indexKastenCounterAll = document.getElementById("indexKastenCounterAll");
  // Prepare http request
  let xhr3 = new XMLHttpRequest();
  xhr3.responseType = `json`;
  url = "http://localhost:8080/karteikasten/";
  xhr3.open("GET", url, true);
  xhr3.setRequestHeader("Content-Type", "application/json");
  // Callback function
  xhr3.onreadystatechange = function() {
    if (xhr3.readyState == 4 && xhr3.status == 200) {
      indexKastenCounterAll.innerHTML = xhr3.response.length;
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
    xhr4.responseType = `json`;
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

// LOGOUT
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

// NEW KARTEI
function newKartei() {
  location.href = "edit.html";
}
