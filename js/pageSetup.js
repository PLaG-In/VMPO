var pageSetup = new function(){
	
	this.hideChatControls = function(){
		$('#chatContainer').hide();
		$('#submitFormContainer').hide();
	}
	
	this.showChatControls = function(){
		var showTime = 500;
		$('#chatContainer').show(showTime);
		$('#submitFormContainer').show(showTime);
	}
};