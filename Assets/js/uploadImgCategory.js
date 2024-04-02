document.addEventListener('DOMContentLoaded', function() {
  const mediaInput = document.getElementById('mediaInput');
  const uploadedImage = document.getElementById('uploadedImage');

  mediaInput.addEventListener('change', function(event) {
    const file = event.target.files[0]; // Get the first file from the selected files
    const reader = new FileReader();

    reader.onload = function(e) {
      // Modifier la propriété CSS background-image de l'élément uploadedImage
      uploadedImage.style.backgroundImage = `url(${e.target.result})`;
      uploadedImage.style.backgroundSize = 'cover'; // Optionnel: ajuster la taille de l'image au conteneur
      uploadedImage.style.backgroundPosition = 'center'; // Optionnel: centrer l'image dans le conteneur
      uploadedImage.style.backgroundColor = 'transparent'; // Changer la couleur de fond en transparent
    };

    if (file) {
      reader.readAsDataURL(file); // Convertir le fichier en URL de données (data URL)
    }
  });
});
