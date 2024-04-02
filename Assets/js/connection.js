const signUpButton = document.getElementById('signUp');
const signInButton = document.getElementById('signIn');
const container = document.getElementById('container');

signUpButton.addEventListener('click', () => {
	container.classList.add("right-panel-active");
});

signInButton.addEventListener('click', () => {
	container.classList.remove("right-panel-active");
});


// Fonction pour extraire la valeur d'un paramètre de requête
function getQueryParam(param) {
  const queryString = window.location.search;
  const urlParams = new URLSearchParams(queryString);
  return urlParams.get(param);
}

const errorValue = getQueryParam('error');
const usernameErrorDiv = document.getElementById('usernameError');
const passwordErrorDiv = document.getElementById('passwordError'); // Sélection de la div
const usernameErrorSignDiv = document.getElementById('usernameErrorSignUp');
const mailErrorDiv = document.getElementById('mailError');

if (errorValue === 'badname') {
  console.log('Nom d\'utilisateur incorrect');
  usernameErrorDiv.textContent = 'Nom d\'utilisateur incorrect'; // Ajout de texte dans la div
  usernameErrorDiv.style.color = 'red';
  usernameErrorDiv.style.fontSize = '11px'; // Modification de la couleur du texte
} else if (errorValue === 'badpassword') {
  console.log('Mot de passe incorrect');
  passwordErrorDiv.textContent = 'Mot de passe incorrect'; // Ajout de texte dans la div
  passwordErrorDiv.style.color = 'red';
  passwordErrorDiv.style.fontSize = '11px'; // Modification de la couleur du texte
  // Autres actions à entreprendre en cas de mot de passe incorrect
} else if (errorValue === 'badnameSignUp') {
  console.log('Nom d\'utilisateur déjà utilisé');
  usernameErrorSignDiv.textContent = 'Nom d\'utilisateur déjà utilisé';
  usernameErrorSignDiv.style.color = 'red';
  usernameErrorSignDiv.style.fontSize = '11px';
  container.classList.add("right-panel-active");
} else if (errorValue === 'badmailSignUp') {
  console.log('Email déjà utilisé');
  mailErrorDiv.textContent = 'Email déjà utilisé';
  mailErrorDiv.style.color = 'red';
  mailErrorDiv.style.fontSize = '11px';
  container.classList.add("right-panel-active");
}

  
