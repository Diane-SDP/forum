document.addEventListener('DOMContentLoaded', () => {
    const searchInput = document.getElementById('searchInput');
    const itemList = document.getElementById('itemList');

    searchInput.addEventListener('input', () => {
        const searchTerm = searchInput.value.toLowerCase();

        for (const li of itemList.getElementsByTagName('li')) {
            if (li.textContent.toLowerCase().includes(searchTerm)) {
                li.style.display = 'block';
            } else {
                li.style.display = 'none';
            }
        }
    });
});