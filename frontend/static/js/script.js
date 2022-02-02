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

function openForm(s) {
  document.getElementById(s).classList.remove("form-disactive");
    document.getElementById(s).classList.add("form-active");
  }
  
  function closeForm(s) {
    document.getElementById(s).classList.add("form-disactive");
    document.getElementById(s).classList.remove("form-active");
  }