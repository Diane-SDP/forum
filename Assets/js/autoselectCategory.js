// Récupérer l'ID de la catégorie à partir de l'URL précédente
function getCategoryFromPreviousUrl() {
    const previousUrl = document.referrer;
    const matches = previousUrl.match(/\/topic\/(\d+)/);
    if (matches && matches.length > 1) {
        return parseInt(matches[1]); // Convertir en nombre entier
    }
    return null;
}

// Sélectionner automatiquement l'option correspondante dans le menu déroulant
function selectCategoryOption(categoryId) {
    const selectElement = document.getElementById('post-category');
    if (!selectElement) return; // Vérifier si l'élément select existe

    for (let option of selectElement.options) {
        if (parseInt(option.value) === categoryId) {
            option.selected = true; // Sélectionner l'option correspondante
            break;
        }
    }
}

// Appel des fonctions pour effectuer la sélection automatique
window.onload = function() {
    const categoryId = getCategoryFromPreviousUrl();
    if (categoryId !== null) {
        selectCategoryOption(categoryId);
    }
};
