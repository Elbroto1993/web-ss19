"use strict";

// Get kastenid from url
let urlParam = function(name, w) {
  w = w || window;
  var rx = new RegExp("[&|?]" + name + "=([^&#]+)"),
    val = w.location.search.match(rx);
  return !val ? "" : val[1];
};
let id = urlParam("Use_Id");

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
  // Get all infos about user
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
  // Get all infos about kasten
  let xhr = new XMLHttpRequest();
  url = `http://localhost:8080/karteikasten/kasten/${id}`;
  xhr.responseType = `json`;
  xhr.open("GET", url, true);
  xhr.onreadystatechange = function() {
    if (xhr.readyState == 4 && xhr.status == 200) {
      let data = xhr.response;
      // Fill elements (name, kategorie, fortschritt) with data from kasten
      document.getElementById("viewName").innerHTML = data.titel;
      document.getElementById("viewKategorie").innerHTML =
        data.ueberkategorie + ">" + data.kategorie;
      document.getElementById("viewFortschritt").innerHTML =
        data.fortschritt + " %";
      // Get username for createdById
      let xhr2 = new XMLHttpRequest();
      let url2 = `http://localhost:8080/user/id/${data.createdByUserId}`;
      xhr2.responseType = `json`;
      xhr2.open("GET", url2, true);
      xhr2.onreadystatechange = function() {
        if (xhr2.readyState == 4 && xhr2.status == 200) {
          let data = xhr2.response;
          // Fill element (erstellt von) with data from createdByUserId
          document.getElementById("viewCreatedBy").innerHTML = data.username;
        }
      };
      xhr2.send(null);
    }
  };
  xhr.send(null);
  // Get all karten pro kasten
  let xhr2 = new XMLHttpRequest();
  let url3 = `http://localhost:8080/karteikarte/kasten/${id}`;
  xhr2.responseType = `json`;
  xhr2.open("GET", url3, true);
  xhr2.onreadystatechange = function() {
    if (xhr2.readyState == 4 && xhr2.status == 200) {
      let data = xhr2.response;
      // Fill element (anzahl karten) with data from length of slice, if data isn't null
      if (!data) {
        document.getElementById("viewCountKarten").innerHTML = 0;
      } else {
        document.getElementById("viewCountKarten").innerHTML = data.length;
      }
      // Loop through all karten and count how many karten are in each fach
      let fach0 = 0;
      let fach1 = 0;
      let fach2 = 0;
      let fach3 = 0;
      let fach4 = 0;
      for (let i = 0; i < data.length; i++) {
        switch (data[i].fach) {
          case "0":
            fach0++;
            break;
          case "1":
            fach1++;
            break;
          case "2":
            fach2++;
            break;
          case "3":
            fach3++;
            break;
          case "4":
            fach4++;
            break;
        }
      }
      document.getElementById("fach0").innerHTML = fach0 + " ";
      document.getElementById("fach1").innerHTML = fach1 + " ";
      document.getElementById("fach2").innerHTML = fach2 + " ";
      document.getElementById("fach3").innerHTML = fach3 + " ";
      document.getElementById("fach4").innerHTML = fach4 + " ";
    }
  };
  xhr2.send(null);
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

// Add eventListener to new kartei button
document.getElementById("newKarteiButton").addEventListener("click", newKartei);
// NEW KARTEI
function newKartei() {
  location.href = "edit.html";
}
