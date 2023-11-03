const effect = document.querySelector(".effect");
const buttons = document.querySelectorAll(".navbar button:not(.plus)");

buttons.forEach((button) => {
  button.addEventListener("click", (e) => {
    const x = e.target.offsetLeft;

    buttons.forEach((btn) => {
      btn.classList.remove("active");
    });

    e.target.classList.add("active");

    anime({
      targets: ".effect",
      left: `${x}px`,
      opacity: "1",
      duration: 600,
    });
  });
});

document.addEventListener("DOMContentLoaded", function () {
  var openModalBtn = document.getElementById("openModalBtn");
  var closeModalBtn = document.getElementById("closeModalBtn");
  var modal = document.getElementById("myModal");

  openModalBtn.addEventListener("click", function () {
    modal.style.display = "block";
  });

  closeModalBtn.addEventListener("click", function () {
    modal.style.display = "none";
  });

  window.addEventListener("click", function (event) {
    if (event.target == modal) {
      modal.style.display = "none";
    }
  });
});

function Num() {
  var modal = document.getElementById("myModal1");
  modal.style.display = "block";
}

function Pass() {
  var modal = document.getElementById("myModal2");
  modal.style.display = "block";
}

function Update() {
  var num = $("#updatenum").val();

  if (num == "") {
    Swal.fire({
      toast: true,
      position: "center",
      icon: "warning",
      title: "Please Complete Required Details",
      showConfirmButton: false,
      timer: 3000,
    });
  } else {
    $.ajax({
      url: "/api/number",
      data: {
        number: num,
        process: "edit",
      },
      success: function () {
        Swal.fire({
          toast: true,
          position: "center",
          icon: "success",
          title: "Update Mobile Number Successful",
          showConfirmButton: false,
          timer: 3000,
        }).then(function () {
          window.location.reload();
        });
      },
    });
  }
}

function Update1() {
  var num = $("#updatepass").val();

  if (num == "") {
    Swal.fire({
      toast: true,
      position: "center",
      icon: "warning",
      title: "Please Complete Required Details",
      showConfirmButton: false,
      timer: 3000,
    });
  } else {
    $.ajax({
      url: "/api/number",
      data: {
        password: num,
        process: "pass",
      },
      success: function () {
        Swal.fire({
          toast: true,
          position: "center",
          icon: "success",
          title: "Update Password Successful",
          showConfirmButton: false,
          timer: 3000,
        }).then(function () {
          window.location.reload();
        });
      },
    });
  }
}

function Cancel() {
  var modal = document.getElementById("myModal1");
  modal.style.display = "none";
}

function Cancel1() {
  var modal = document.getElementById("myModal2");
  modal.style.display = "none";
}
