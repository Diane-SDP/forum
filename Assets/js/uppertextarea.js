const textarea = document.getElementById('contentInput');

textarea.addEventListener('input', function() {
    textarea.style.height = 'auto';
    textarea.style.height = Math.min(textarea.scrollHeight, 200) + 'px'; 
});
