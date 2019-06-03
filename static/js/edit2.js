"use strict";

// MDE EDITOR
let simplemde = new SimpleMDE({
  element: document.getElementById("edit2Textarea1")
});

let simplemde2 = new SimpleMDE({
  element: document.getElementById("edit2Textarea2")
});

// Get kastenid from url
let urlParam = function(name, w) {
  w = w || window;
  var rx = new RegExp("[&|?]" + name + "=([^&#]+)"),
    val = w.location.search.match(rx);
  return !val ? "" : val[1];
};
let kastenid = urlParam("_kastenid");

// Get karteid from url
let urlParam2 = function(name, w) {
  w = w || window;
  var rx = new RegExp("[&|?]" + name + "=([^&#]+)"),
    val = w.location.search.match(rx);
  return !val ? "" : val[1];
};
let karteid = urlParam("_karteid");

// Functions for buttons when no karte is selected
document.getElementById("saveButton").addEventListener("click", saveKarte);
function saveKarte() {
  // Get values from single card forms
  let titel = document.getElementById("edit2input").value;
  let frage = simplemde.value();
  let antwort = simplemde2.value();
  // Prepare http post request
  let xhr = new XMLHttpRequest();
  let url = `http://localhost:8080/add-or-update-karte`;
  xhr.open("POST", url, true);
  xhr.onreadystatechange = function() {
    if (xhr.readyState == 4 && xhr.status == 200) {
      let karteid = xhr.response.replace(/"/g, "");
      window.location.href = `http://localhost:8080/edit2?_kastenid=${kastenid}`;
    }
  };
  let data = JSON.stringify({
    _id: karteid,
    kastenid: kastenid,
    titel: titel,
    frage: frage,
    antwort: antwort,
    fach: "0"
  });
  xhr.send(data);
}

// function deleteKarte() {
//   let id = this.parentNode.id;
//   // Toggle modal for second delete button
//   document.getElementById("profilDeleteModal").classList.add("is-active");

//   document
//     .getElementById("modal-close")
//     .addEventListener("click", untoggleDeleteModal);
//   document
//     .getElementById("profilKeepButton")
//     .addEventListener("click", untoggleDeleteModal);

//   function untoggleDeleteModal() {
//     document.getElementById("profilDeleteModal").classList.remove("is-active");
//   }
//   // Final delete button
//   document
//     .getElementById("profilFinallyedit2deleteButton")
//     .addEventListener("click", deleteKarte);
//   function deleteKarte() {
//     // Delete kasten
//     let xhr = new XMLHttpRequest();
//     let url = `http://localhost:8080/karteikarte/${id}`;
//     xhr.open("DELETE", url, true);
//     xhr.onreadystatechange = function() {
//       if (xhr.readyState == 4 && xhr.status == 200) {
//         location.reload();
//       }
//     };
//     xhr.send(null);
//   }
// }
