//pageSetup.showLoginControls();
var pageSetup = new function(){
	
	this.hideLoginControls = function(){
		$('#loginBtn').hide();
		$('#signUpBtn').hide();
		$('#exitBtn').show();
		$('#userName').show();
	}
	
	this.showLoginControls = function(){
		var showTime = 500;
		$('#loginBtn').show();
		$('#signUpBtn').show();
		$('#exitBtn').hide();
		$('#userName').hide();
	}
};