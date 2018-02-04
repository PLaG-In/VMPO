var SERVER_URL = "http://localhost:8080";

function loginRequest(){
		var url = SERVER_URL + '/auth';

 		var resultLabel = $('#loginResult');
		resultLabel.text("");

		var userStr = $('#alogin').val();
		var passwordStr = $('#apassword').val();
		if (userStr == '' || passwordStr == '')
		{
			alert('Please fill all fields');
			return false;
		}
		pageSetup.hideLoginControls();
		$('#userName').text(userStr);

		var postData = { login : userStr, password : passwordStr };
		$.ajax({
			type: 'POST',
			url: url,
			data: postData,
			dataType: 'json',
			success: function(result){
				if (result.status != "OK")
				{
					switch (result.message)
					{
						case "WRONG_PASSWORD":
							resultLabel.text("Wrong password");
							break;
						case "WRONG_LOGIN":
							resultLabel.text("Wrong login");
							break;
						case "ALREADY_LOGGED":
							resultLabel.text("User has already logged in");
							break;
						default:
							resultLabel.text("Unknown server response");
							break;
					}
				}
	 			else
				{
			 		pageSetup.showLoginControls();
				}
			},
			error: function(){
				console.log("[LOGIN] Unhandled server error");
			}
		});
	return true;
}

function exitRequest(){
	var url = SERVER_URL + '/sign_out';
	var postData = { login : userStr, command : 'remove_visitor'};
	console.log(postData);
	pageSetup.showLoginControls();
	return $.ajax({
		async: false,
		type: 'POST',
		url: url,
		data: postData,
		dataType: 'json',
		error: function(result){
			console.log("[REMOVE_VISITOR] Unhandled server error");
		}
	});
}

function signupRequest(){
		var url = SERVER_URL + '/reg';

 		var resultLabel = $('#loginResult');
		resultLabel.text("");

		var userStr = $('#slogin').val();
		var passwordStr = $('#spassword').val();
		var repasswordStr = $('#repassword').val();
		if (userStr == '' || passwordStr == '')
		{
			alert('Заполните все поля');
			return false;
		}

		if (passwordStr != repasswordStr)
		{
			alert('Пароли не совпадают');
			return false;
		}
		pageSetup.hideLoginControls();
		$('#userName').text(userStr);
		var postData = { login : userStr, password : passwordStr };
		$.ajax({
			type: 'POST',
			url: url,
			data: postData,
			dataType: 'json',
			success: function(result){
				if (result.status != "OK")
				{
					switch (result.message)
					{
						case "WRONG_PASSWORD":
							resultLabel.text("Wrong password");
							break;
						case "WRONG_LOGIN":
							resultLabel.text("Wrong login");
							break;
						case "ALREADY_LOGGED":
							resultLabel.text("User has already logged in");
							break;
						default:
							resultLabel.text("Unknown server response");
							break;
					}
				}
				else
				{
			 		pageSetup.showLoginControls();
				}
			},
			error: function(){
				console.log("[SIGNUP] Unhandled server error");
			}
		});
	return true;
}
