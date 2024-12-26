const IMAGE_EXTENSIONS = ["jpg", "jpeg", "png", "gif"];
const VIDEO_EXTENSIONS = ["mp4", "avi", "mov", "wmv"];
let currentFileType = null;
let currentExtension = null;

async function checkFileExists(type, filename, extensions) {
  for (const ext of extensions) {
    try {
      const response = await fetch(`/download/${type}/${filename}.${ext}`);
      if (response.ok) {
        return { exists: true, extension: ext };
      }
    } catch (error) {
      console.log(`Failed to check .${ext} file:`, error);
    }
  }
  return { exists: false };
}

async function loadPreview() {
  const pathParts = window.location.pathname.split("/");
  const filename = pathParts[pathParts.length - 1];

  const previewImage = document.getElementById("previewImage");
  const previewVideo = document.getElementById("previewVideo");
  const previewError = document.getElementById("previewError");
  const loadingSpinner = document.getElementById("loadingSpinner");
  const photoBtn = document.getElementById("photoBtn");
  const videoBtn = document.getElementById("videoBtn");
  const previewMessage = document.getElementById("previewMessage");

  // 초기 상태 설정
  previewImage.style.display = "none";
  previewImage.classList.remove("loaded");
  previewVideo.style.display = "none";
  previewVideo.classList.remove("loaded");
  previewError.style.display = "none";
  loadingSpinner.style.display = "flex";
  photoBtn.disabled = true;
  videoBtn.disabled = true;
  photoBtn.classList.remove("active");
  videoBtn.classList.remove("active");
  previewMessage.textContent = "";

  // 이미지 체크
  const imageResult = await checkFileExists(
    "images",
    filename,
    IMAGE_EXTENSIONS
  );
  if (imageResult.exists) {
    previewImage.src = `/download/images/${filename}.${imageResult.extension}`;
    previewImage.onload = () => {
      loadingSpinner.style.display = "none";
      previewImage.style.display = "block";
      // 약간의 지연 후 페이드인 효과 적용
      setTimeout(() => {
        previewImage.classList.add("loaded");
      }, 50);
      photoBtn.disabled = false;
      photoBtn.classList.add("active");
      previewMessage.textContent = "사진이 준비되었습니다.";
      currentFileType = "images";
      currentExtension = imageResult.extension;
    };
    previewImage.onerror = showError;
    return;
  }

  // 비디오 체크
  const videoResult = await checkFileExists(
    "videos",
    filename,
    VIDEO_EXTENSIONS
  );
  if (videoResult.exists) {
    previewVideo.src = `/download/videos/${filename}.${videoResult.extension}`;
    previewVideo.style.display = "block";
    // 약간의 지연 후 페이드인 효과 적용
    setTimeout(() => {
      previewVideo.classList.add("loaded");
    }, 50);
    loadingSpinner.style.display = "none";
    videoBtn.disabled = false;
    videoBtn.classList.add("active");
    previewMessage.textContent = "동영상이 준비되었습니다 💝";
    currentFileType = "videos";
    currentExtension = videoResult.extension;
    return;
  }

  showError();
}

function showError() {
  const previewError = document.getElementById("previewError");
  const loadingSpinner = document.getElementById("loadingSpinner");
  const previewMessage = document.getElementById("previewMessage");

  loadingSpinner.style.display = "none";
  previewError.style.display = "block";
  previewMessage.textContent = "파일을 찾을 수 없습니다 😢";
}

async function downloadFile(type) {
  if (currentFileType !== type) return;

  const pathParts = window.location.pathname.split("/");
  const filename = pathParts[pathParts.length - 1];

  const a = document.createElement("a");
  a.href = `/download/${type}/${filename}.${currentExtension}`;
  a.download = `${filename}.${currentExtension}`;
  document.body.appendChild(a);
  a.click();
  document.body.removeChild(a);
}

// 페이지 로드 시 미리보기 시작
window.onload = loadPreview;
