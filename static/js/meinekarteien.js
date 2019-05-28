"use strict";

let lernButtons = document.getElementsByClassName("kastenLernButton");
for (let i = 0; i < lernButtons.length; i++) {
  lernButtons[i].addEventListener("click", function() {
    let id = lernButtons[i].parentNode.id;
    window.location.href = `http://localhost:8080/lern?_kastenid=${id}`;
  });
}

let editButtons = document.getElementsByClassName("kastenEditButton");
for (let i = 0; i < editButtons.length; i++) {
  editButtons[i].addEventListener("click", function() {
    let id = editButtons[i].parentNode.id;
    window.location.href = `http://localhost:8080/edit2?_kastenid=${id}`;
  });
}

let deleteButtons = document.getElementsByClassName("kastenDeleteButton");
for (let i = 0; i < deleteButtons.length; i++) {
  deleteButtons[i].addEventListener("click", deleteKasten);
}

function deleteKasten() {
  let id = this.parentNode.id;
  // Toggle modal for second delete button
  document.getElementById("profilDeleteModal").classList.add("is-active");

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
  function deleteProfile() {}
}
