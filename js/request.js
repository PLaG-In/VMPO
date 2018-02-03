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
		
		var postData = { user : userStr, password : passwordStr };
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
			/* 		pageSetup.showChatControls();
					$('#apassword').val('');
					chat.username = userStr;
					var chatUpdateInterval = setInterval(chat.updateAll, 200);
					$(window).on('unload', function(){
						clearInterval(chatUpdateInterval);
						return chat.exit();
					}); */
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
	var postData = { user : chat.username, command : 'remove_visitor'};
	console.log(postData);
		
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
		
		var postData = { user : userStr, password : passwordStr };
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
			/* 		pageSetup.showChatControls();
					$('#apassword').val('');
					chat.username = userStr;
					var chatUpdateInterval = setInterval(chat.updateAll, 200);
					$(window).on('unload', function(){
						clearInterval(chatUpdateInterval);
						return chat.exit();
					}); */
				}
			},
			error: function(){
				console.log("[SIGNUP] Unhandled server error");
			}
		});
	return true;
}