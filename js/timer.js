/* function timer(_time, _call){
			timer.lastCall = _call
			timer.lastTime = _time
			timer.timerInterval = setInterval(function(){
						_call(_time[0],_time[1],_time[2]);
						_time[2]--
						if(_time[0]==0 && _time[1]==0 && _time[2]==0){
									timer.pause()
									_call(0,0,0);
						}
						if(_time[2]==0){
									_time[2] = 59
									_time[1]--
									if(_time[1]==0){
												_time[1] = 59
												_time[0]--
									}
						}
						timer.lastTime = _time
			}, 1000)
} */
/* timer.pause = function(){
	clearInterval(timer.timerInterval)
}
timer.start = function(){
	timer([0,0,30], function(h,m,s){
		document.getElementById("timer").innerHTML = h + ':' +m + ':' +s
	});
	timer(timer.lastTime, timer.lastCall)
} */
var pause = false;	
function startTimer() {
	pause = false;
    var my_timer = document.getElementById("time");
    var time = my_timer.innerHTML;
    var arr = time.split(":");
    var h = arr[0];
    var m = arr[1];
    var s = arr[2];
	if (pause) {
		return;
	}
    if (s == 0) {
      if (m == 0) {
        if (h == 0) {
          alert("Время вышло");
          //window.location.reload();
          return;
        }
        h--;
        m = 60;
        if (h < 10) h = "0" + h;
      }
      m--;
      if (m < 10) m = "0" + m;
      s = 59;
    }
    else s--;
    if (s < 10) s = "0" + s;
    document.getElementById("time").innerHTML = h+":"+m+":"+s;
    setTimeout(startTimer, 1000);
}

function stopTimer() {
	pause = true;
}
