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

// Get kastenid from url
let urlParam = function(name, w) {
  w = w || window;
  var rx = new RegExp("[&|?]" + name + "=([^&#]+)"),
    val = w.location.search.match(rx);
  return !val ? "" : val[1];
};
let id = urlParam("Use_Id");

// STORE NEW KASTEN
document
  .getElementById("editKastenButton")
  .addEventListener("click", newKasten);

function newKasten(e) {
  e.preventDefault();
  // Get data from forms
  let kategorieError = document.getElementById("kategoriePflicht");
  let titel = document.getElementById("editKastenTitel").value;
  let kategorieHelp = document.getElementById("editKastenSelect");
  let kategorie = kategorieHelp.options[kategorieHelp.selectedIndex].value;
  let beschreibung = document.getElementById("editKastenBeschreibung").value;
  let ueberkategorie = kategorieHelp.options[
    kategorieHelp.selectedIndex
  ].parentNode.getAttribute("label");
  // Clear errors
  kategorieError.innerHTML = "";
  kategorieError.classList.remove("registerAlert");
  // Get value from radio buttons
  let privateHelp = document.getElementsByClassName("editRadioButton");
  let privateOrPublic;
  for (let i = 0; i < privateHelp.length; i++) {
    if (privateHelp[i].checked) {
      privateOrPublic = privateHelp[i].value;
    }
  }
  // Change value to boolean
  if (privateOrPublic === "Öffentlich") {
    privateOrPublic = "false";
  } else {
    privateOrPublic = "true";
  }
  // Check if kategorie is selected
  if (!kategorie) {
    kategorieError.innerHTML = "Pflichtfeld!";
    kategorieError.classList.add("registerAlert");
  } else {
    // If url has id update kasten, else store new kasten
    if (id) {
      let xhr = new XMLHttpRequest();
      let url = "http://localhost:8080/karteikasten/update";
      xhr.open("POST", url, true);
      xhr.setRequestHeader("Content-Type", "application/json");
      xhr.onreadystatechange = function() {
        if (xhr.readyState == 4 && xhr.status == 200) {
          console.log(xhr.responseText);
        }
      };
      // Set data for POST
      let data = JSON.stringify({
        kastenid: id,
        titel: titel,
        kategorie: kategorie,
        beschreibung: beschreibung,
        private: privateOrPublic,
        ueberkategorie: ueberkategorie
      });
      xhr.send(data);
      // Clear forms
      document.getElementById("editKastenTitel").value = "";
      document.getElementById("editKastenBeschreibung").value = "";
    } else {
      console.log("id nicht vorhan");
      let xhr = new XMLHttpRequest();
      let url = "http://localhost:8080/karteikasten/";
      xhr.open("POST", url, true);
      xhr.setRequestHeader("Content-Type", "application/json");
      xhr.onreadystatechange = function() {
        if (xhr.readyState == 4 && xhr.status == 200) {
          console.log(xhr.responseText);
        }
      };
      // Set data for POST
      let data = JSON.stringify({
        titel: titel,
        kategorie: kategorie,
        beschreibung: beschreibung,
        private: privateOrPublic,
        ueberkategorie: ueberkategorie
      });
      xhr.send(data);
      // Clear forms
      document.getElementById("editKastenTitel").value = "";
      document.getElementById("editKastenBeschreibung").value = "";
    }
  }
}

// Add eventListener to new kartei button
document.getElementById("newKarteiButton").addEventListener("click", newKartei);
// NEW KARTEI
function newKartei() {
  location.href = "edit.html";
}
