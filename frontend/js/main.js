(function () {
  "use strict";

  // Мобильное меню
  var burger = document.querySelector(".header__burger");
  var mobileNav = document.getElementById("mobileNav");
  var mobileLinks = document.querySelectorAll(".header__mobile-list a");

  if (burger && mobileNav) {
    burger.addEventListener("click", function () {
      var isOpen = mobileNav.classList.toggle("is-open");
      burger.classList.toggle("is-open", isOpen);
      burger.setAttribute("aria-expanded", isOpen);
      mobileNav.setAttribute("aria-hidden", !isOpen);
      document.body.style.overflow = isOpen ? "hidden" : "";
    });

    mobileLinks.forEach(function (link) {
      link.addEventListener("click", function () {
        mobileNav.classList.remove("is-open");
        burger.classList.remove("is-open");
        burger.setAttribute("aria-expanded", "false");
        mobileNav.setAttribute("aria-hidden", "true");
        document.body.style.overflow = "";
      });
    });
  }

  // Плавная прокрутка для якорных ссылок
  document.querySelectorAll('a[href^="#"]').forEach(function (anchor) {
    anchor.addEventListener("click", function (e) {
      var href = this.getAttribute("href");
      if (href === "#") return;
      var target = document.querySelector(href);
      if (target) {
        e.preventDefault();
        target.scrollIntoView({ behavior: "smooth", block: "start" });
      }
    });
  });

  // Кнопка «Вернуться наверх»
  var backToTop = document.getElementById("backToTop");
  if (backToTop) {
    backToTop.addEventListener("click", function (e) {
      e.preventDefault();
      window.scrollTo({ top: 0, behavior: "smooth" });
    });
  }

  // Анимация появления при скролле
  var observerOptions = {
    root: null,
    rootMargin: "0px 0px -60px 0px",
    threshold: 0.1,
  };
  var observer = new IntersectionObserver(function (entries) {
    entries.forEach(function (entry) {
      if (entry.isIntersecting) {
        entry.target.classList.add("is-visible");
      }
    });
  }, observerOptions);

  document.querySelectorAll(".animate-on-scroll").forEach(function (el) {
    observer.observe(el);
  });

  // Форма обратной связи (fetch)
  var form = document.getElementById("contactForm");
  var formSuccess = document.getElementById("formSuccess");
  var formError = document.getElementById("formError");
  var formSubmitBtn = document.getElementById("formSubmit");

  if (form) {
    form.addEventListener("submit", function (e) {
      e.preventDefault();

      formSuccess.classList.remove("is-visible");
      formError.classList.remove("is-visible");
      formError.textContent = "";

      var name = document.getElementById("formName").value.trim();
      var phone = document.getElementById("formPhone").value.trim();

      if (!name || !phone) {
        formError.textContent = "Заполните имя и телефон.";
        formError.classList.add("is-visible");
        return;
      }

      if (formSubmitBtn) {
        formSubmitBtn.disabled = true;
        formSubmitBtn.textContent = "Отправка…";
      }

      // Всегда используем полный URL к бэкенду
      // В Docker окружении фронтенд на порту 8000, бэкенд на 8080
      var apiUrl =
        form.getAttribute("data-action") || "https://ecolife.ru/api/feedback";

      var controller = new AbortController();
      var timeoutId = setTimeout(function () {
        controller.abort();
      }, 10000);

      fetch(apiUrl, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ name: name, phone: phone }),
        signal: controller.signal,
      })
        .then(function (res) {
          clearTimeout(timeoutId);
          if (res.ok) {
            formSuccess.classList.add("is-visible");
            form.reset();
          } else {
            return res
              .json()
              .then(function (data) {
                throw { response: { data: data } };
              })
              .catch(function (err) {
                if (err.response) throw err;
                throw new Error(res.statusText || "Ошибка сервера");
              });
          }
        })
        .catch(function (err) {
          clearTimeout(timeoutId);
          var message =
            "Не удалось отправить заявку. Попробуйте позже или позвоните нам.";
          if (err.response && err.response.data && err.response.data.message) {
            message = err.response.data.message;
          } else if (
            err.name === "AbortError" ||
            err.message === "Failed to fetch"
          ) {
            message = "Нет связи с сервером. Позвоните нам: +7 928 007-38-38";
          }
          formError.textContent = message;
          formError.classList.add("is-visible");
        })
        .finally(function () {
          if (formSubmitBtn) {
            formSubmitBtn.disabled = false;
            formSubmitBtn.textContent = "Связаться";
          }
        });
    });
  }
})();
