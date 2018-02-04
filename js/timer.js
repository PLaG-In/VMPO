var timerPaused = false;

function startTimer() {
	if (timerPaused) return;
	var my_timer = document.getElementById("time");
	var time = my_timer.innerHTML;
	var arr = time.split(":");
	var h = arr[0];
	var m = arr[1];
	var s = arr[2];
	if (s == 59) {
		if (m == 59)
		{
			h++;
			m = 0;
			if (h < 10) h = "0" + h;
		}
		else m++;
		if (m < 10) m = "0" + m;
		s = 0;
	}
	else s++;
	if (s < 10) s = "0" + s;
	document.getElementById("time").innerHTML = h+":"+m+":"+s;
	setTimeout(startTimer, 1000);
 }

 function getTime(){
	 return document.getElementById("time").innerHTML;
 }

$("#timer-button").on("click", function() {
	if($(this).attr("play") == 1) {
		if (isSelected()) {
				timerPaused = false;
				$(this).attr("play", "0");
				chooseTask();
				document.getElementById("time").innerHTML = getTaskTime();
				startTimer();
		} else {
			alert("Для запуска таймера выберите задачу!");
		}
	} else {
        return;
    }
});

$("#stop-button").on("click", function() {
    if($("#timer-button").attr("play") == 0) {
        $("#timer-button").attr("play", "1");
				timerPaused = true;
				putTime();
				document.getElementById("time").innerHTML = "00:00:00";
    }
    else {
        return;
    }
});
