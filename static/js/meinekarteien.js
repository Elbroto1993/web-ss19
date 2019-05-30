"use strict";

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
