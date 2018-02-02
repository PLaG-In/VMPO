var SERVER_URL = "http://localhost:8080/";
		// 1. Создаём новый объект XMLHttpRequest
		// Get the modal
	var signup = document.getElementById('id01');
	var login = document.getElementById('id02');

	function loadPhones() {
		var xhr = new XMLHttpRequest();

		xhr.open('GET', SERVER_URL + "auth", true, 'sds', 'sdfsd');
		xhr.setRequestHeader('Content-Type');
		xhr.setRequestHeader('login', '123');
		xhr.setRequestHeader('password', '456');
	  
		xhr.send();

		if (xhr.status != 200) {
			// обработать ошибку
			alert('Ошибка ' + xhr.status + ': ' + xhr.statusText);
		} else {
			// вывести результат
			alert(xhr.responseText);
		}
    }

	// When the user clicks anywhere outside of the modal, close it
	window.onclick = function(event) {
		if (event.target == signup) {
			signup.style.display = "none";
		} else if (event.target == login){
			login.style.display = "none";
		}
	}
	
    function Calendar2(id, year, month) {   
    	var Dlast = new Date(year,month+1,0).getDate(),
    	    D = new Date(year,month,Dlast),
    	    DNlast = new Date(D.getFullYear(),D.getMonth(),Dlast).getDay(),
    	    DNfirst = new Date(D.getFullYear(),D.getMonth(),1).getDay(),
    	    calendar = '<tr>',
    	    month=["Январь","Февраль","Март","Апрель","Май","Июнь","Июль","Август","Сентябрь","Октябрь","Ноябрь","Декабрь"];
    	if (DNfirst != 0) 
    	{
    		for(var  i = 1; i < DNfirst; i++) calendar += '<td>';
    	}
    	else
    	{
      		for(var  i = 0; i < 6; i++) calendar += '<td>';
    	}

    	for(var  i = 1; i <= Dlast; i++) {
      		if (i == new Date().getDate() && D.getFullYear() == new Date().getFullYear() && D.getMonth() == new Date().getMonth()) 
      		{
        		calendar += '<td class="cell today">' + i;
      		}
      		else
      		{
        		calendar += '<td class = "cell">' + i;
      		}
      		if (new Date(D.getFullYear(),D.getMonth(),i).getDay() == 0) 
      		{
        		calendar += '<tr>';
      		}
    	}
    	for(var  i = DNlast; i < 7; i++) calendar += '<td>&nbsp;';
    	document.querySelector('#'+id+' tbody').innerHTML = calendar;
    	document.querySelector('#'+id+' thead td:nth-child(2)').innerHTML = month[D.getMonth()] +' '+ D.getFullYear();
    	document.querySelector('#'+id+' thead td:nth-child(2)').dataset.month = D.getMonth();
    	document.querySelector('#'+id+' thead td:nth-child(2)').dataset.year = D.getFullYear();
    	if (document.querySelectorAll('#'+id+' tbody tr').length < 6) 
    	{  // чтобы при перелистывании месяцев не "подпрыгивала" вся страница, добавляется ряд пустых клеток. Итог: всегда 6 строк для цифр
        	document.querySelector('#'+id+' tbody').innerHTML += '<tr><td>&nbsp;<td>&nbsp;<td>&nbsp;<td>&nbsp;<td>&nbsp;<td>&nbsp;<td>&nbsp;';
    	}
    }

    Calendar2("calendar2", new Date().getFullYear(), new Date().getMonth());
    // переключатель минус месяц
    document.querySelector('#calendar2 thead tr:nth-child(1) td:nth-child(1)').onclick = function() {
      Calendar2("calendar2", document.querySelector('#calendar2 thead td:nth-child(2)').dataset.year, parseFloat(document.querySelector('#calendar2 thead td:nth-child(2)').dataset.month)-1);      
    }
    // переключатель плюс месяц
    document.querySelector('#calendar2 thead tr:nth-child(1) td:nth-child(3)').onclick = function() {
      Calendar2("calendar2", document.querySelector('#calendar2 thead td:nth-child(2)').dataset.year, parseFloat(document.querySelector('#calendar2 thead td:nth-child(2)').dataset.month)+1);
    }

	$(document).ready(function() 
	{			
	    $('table').on('click', '.cell', function () 
	    {	
	        var index = $('.cell').index();	
	        $('.cell').filter(function(){
	        	return $(this).index() != index;
	        }).removeClass('click');
	    	
	    	$(this).toggleClass('click');
	    });
	});
