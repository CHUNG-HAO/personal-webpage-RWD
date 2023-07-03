window.addEventListener('DOMContentLoaded', () => {
    const mask = document.querySelector('.mask');
    const dynamicText = document.querySelector('.dynamic-text');
    const mainContent = document.querySelector('.main-content');
    dynamicText.textContent = 'Chris Chung 個人網頁';
  
    dynamicText.addEventListener('animationend', () => {
      mask.style.display = 'none';
      mainContent.style.opacity = 1;
    });
  });
  
  