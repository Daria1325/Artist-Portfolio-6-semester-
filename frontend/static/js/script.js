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

function openForm(id) {
  //ссылка которая сейчас + "delete/"+ id
  //вставить в кнопку делете из формы ссылку <a href=s>DELETE</a>
  // document.getElementById(s).classList.remove("form-disactive");
  //   document.getElementById(s).classList.add("form-active");
  }
  
  function closeForm(s) {
    document.getElementById(s).classList.add("form-disactive");
    document.getElementById(s).classList.remove("form-active");
  }