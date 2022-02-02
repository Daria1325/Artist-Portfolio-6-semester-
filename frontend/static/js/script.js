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
    document.getElementById("form_edit").style.display = "block";
  }
  
  function closeForm() {
    document.getElementById("form_edit").style.display = "none";
  }