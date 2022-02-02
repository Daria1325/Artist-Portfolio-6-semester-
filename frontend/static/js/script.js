var expanded = false;

function showCheckboxes() {
  var checkboxes = document.getElementById("checkboxes");
  if (!expanded) {
    checkboxes.style.display = "block";
    expanded = true;
  } else {
    checkboxes.style.display = "none";
    expanded = false;
  }
}

function openEditForm() {
  document.getElementById("form_edit").classList.remove("form-disactive");
    document.getElementById("form_edit").classList.add("form-active");
  }
  
  function closeForm() {
    document.getElementById("form_edit").classList.add("form-disactive");
    document.getElementById("form_edit").classList.remove("form-active");
  }