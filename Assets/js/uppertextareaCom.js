const textarea = document.getElementById('commentInput');

textarea.addEventListener('input', function() {
    textarea.style.height = 'auto';
    textarea.style.height = Math.min(textarea.scrollHeight, 80) + 'px'; 
});
