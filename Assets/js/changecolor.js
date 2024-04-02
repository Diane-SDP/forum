document.addEventListener('DOMContentLoaded', function() {
    var colorChoice = localStorage.getItem('colorChoice');
    if (colorChoice) {
        var menu = document.getElementById('menuboard');
        menu.style.backgroundColor = colorChoice;
    }
});


