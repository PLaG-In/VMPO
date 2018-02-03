		// Get the modal
var signup = document.getElementById('id01');
var login = document.getElementById('id02');

	// When the user clicks anywhere outside of the modal, close it
window.onclick = function(event) {
	if (event.target == signup) {
		signup.style.display = "none";
	} else if (event.target == login){
		login.style.display = "none";
	}
}