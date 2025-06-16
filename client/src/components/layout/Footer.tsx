import { Link } from "react-router-dom";

const Footer = () => {
  return (
    <footer className="bg-conecta-blue-dark text-white py-10">
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div className="grid grid-cols-1 md:grid-cols-4 gap-8">
          <div>
            <h3 className="text-xl font-bold mb-4">ConectaMaré</h3>
            <p className="text-gray-300">
              Conectando clientes a profissionais qualificados de forma rápida,
              segura e eficiente.
            </p>
          </div>
          <div>
            <h4 className="font-semibold mb-4">Para Clientes</h4>
            <ul className="space-y-2 text-gray-300">
              <li>
                <Link to="/professionals" className="hover:text-white">
                  Buscar Profissionais
                </Link>
              </li>
              <li>
                <Link to="/how-it-works" className="hover:text-white">
                  Como Funciona
                </Link>
              </li>
              <li>
                <Link to="/pricing" className="hover:text-white">
                  Preços
                </Link>
              </li>
              <li>
                <Link to="/reviews" className="hover:text-white">
                  Avaliações
                </Link>
              </li>
            </ul>
          </div>
          <div>
            <h4 className="font-semibold mb-4">Para Profissionais</h4>
            <ul className="space-y-2 text-gray-300">
              <li>
                <Link to="/forpros" className="hover:text-white">
                  Por que se Cadastrar
                </Link>
              </li>
              <li>
                <Link to="/success-stories" className="hover:text-white">
                  Histórias de Sucesso
                </Link>
              </li>
              <li>
                <Link to="/pro-resources" className="hover:text-white">
                  Recursos
                </Link>
              </li>
              <li>
                <Link to="/pro-pricing" className="hover:text-white">
                  Planos
                </Link>
              </li>
            </ul>
          </div>
          <div>
            <h4 className="font-semibold mb-4">Empresa</h4>
            <ul className="space-y-2 text-gray-300">
              <li>
                <Link to="/about" className="hover:text-white">
                  Sobre nós
                </Link>
              </li>
              <li>
                <Link to="/contact" className="hover:text-white">
                  Contato
                </Link>
              </li>
              <li>
                <Link to="/privacy" className="hover:text-white">
                  Privacidade
                </Link>
              </li>
              <li>
                <Link to="/terms" className="hover:text-white">
                  Termos de Uso
                </Link>
              </li>
            </ul>
          </div>
        </div>
        <div className="border-t border-gray-600 mt-8 pt-8 flex flex-col md:flex-row justify-between items-center">
          <p>
            © {new Date().getFullYear()} ConectaMaré. Todos os direitos
            reservados.
          </p>
          <div className="flex space-x-4 mt-4 md:mt-0">
            <a href="#" className="text-gray-300 hover:text-white">
              <span className="sr-only">Facebook</span>
              <svg className="h-6 w-6" fill="currentColor" viewBox="0 0 24 24">
                <path
                  fillRule="evenodd"
                  clipRule="evenodd"
                  d="M22 12c0-5.523-4.477-10-10-10S2 6.477 2 12c0 4.991 3.657 9.128 8.438 9.878v-6.987h-2.54V12h2.54V9.797c0-2.506 1.492-3.89 3.777-3.89 1.094 0 2.238.195 2.238.195v2.46h-1.26c-1.243 0-1.63.771-1.63 1.562V12h2.773l-.443 2.89h-2.33v6.988C18.343 21.128 22 16.991 22 12z"
                />
              </svg>
            </a>
            <a href="#" className="text-gray-300 hover:text-white">
              <span className="sr-only">Instagram</span>
              <svg className="h-6 w-6" fill="currentColor" viewBox="0 0 24 24">
                <path
                  fillRule="evenodd"
                  clipRule="evenodd"
                  d="M12.315 2c2.43 0 2.784.013 3.808.06 1.064.049 1.791.218 2.427.465.668.25 1.221.609 1.784 1.172.56.56.921 1.11 1.172 1.784.247.636.416 1.363.465 2.427.048 1.067.06 1.407.06 4.123v.08c0 2.643-.012 2.987-.06 4.043-.049 1.064-.218 1.791-.465 2.427a4.903 4.903 0 01-1.172 1.784c-.56.56-1.11.921-1.784 1.172-.636.247-1.363.416-2.427.465-1.067.048-1.407.06-4.123.06h-.08c-2.643 0-2.987-.012-4.043-.06-1.064-.049-1.791-.218-2.427-.465a4.903 4.903 0 01-1.784-1.172 4.903 4.903 0 01-1.172-1.784c-.247-.636-.416-1.363-.465-2.427-.047-1.024-.06-1.379-.06-3.808v-.63c0-2.43.013-2.784.06-3.808.049-1.064.218-1.791.465-2.427a4.903 4.903 0 011.172-1.784A4.903 4.903 0 016.455 2.525c.636-.247 1.363-.416 2.427-.465C9.906 2.013 10.26 2 12.315 2zm0 1.802h-.63c-2.506 0-2.784.011-3.807.058-.917.042-1.425.196-1.758.325-.433.168-.741.37-1.05.683-.316.31-.518.618-.683 1.05-.13.333-.284.842-.325 1.758-.047 1.023-.058 1.351-.058 3.807v.63c0 2.506.011 2.784.058 3.807.042.917.196 1.425.325 1.758.168.433.37.741.683 1.05.31.316.618.518 1.05.683.333.13.842.284 1.758.325 1.054.048 1.37.058 4.041.058h.08c2.597 0 2.917-.01 3.96-.058.976-.045 1.505-.207 1.858-.344.466-.182.8-.398 1.15-.748.35-.35.566-.683.748-1.15.137-.353.3-.882.344-1.857.048-1.055.058-1.37.058-4.041v-.08c0-2.597-.01-2.917-.058-3.96-.045-.976-.207-1.505-.344-1.858a3.09 3.09 0 00-.748-1.15 3.09 3.09 0 00-1.15-.748c-.353-.137-.882-.3-1.857-.344-1.023-.047-1.351-.058-3.807-.058zM12 6.865a5.135 5.135 0 110 10.27 5.135 5.135 0 010-10.27zm0 1.802a3.333 3.333 0 100 6.666 3.333 3.333 0 000-6.666zm5.338-3.205a1.2 1.2 0 110 2.4 1.2 1.2 0 010-2.4z"
                />
              </svg>
            </a>
            <a href="#" className="text-gray-300 hover:text-white">
              <span className="sr-only">Twitter</span>
              <svg className="h-6 w-6" fill="currentColor" viewBox="0 0 24 24">
                <path d="M8.29 20.251c7.547 0 11.675-6.253 11.675-11.675 0-.178 0-.355-.012-.53A8.348 8.348 0 0022 5.92a8.19 8.19 0 01-2.357.646 4.118 4.118 0 001.804-2.27 8.224 8.224 0 01-2.605.996 4.107 4.107 0 00-6.993 3.743 11.65 11.65 0 01-8.457-4.287 4.106 4.106 0 001.27 5.477A4.072 4.072 0 012.8 9.713v.052a4.105 4.105 0 003.292 4.022 4.095 4.095 0 01-1.853.07 4.108 4.108 0 003.834 2.85A8.233 8.233 0 012 18.407a11.616 11.616 0 006.29 1.84" />
              </svg>
            </a>
          </div>
        </div>
      </div>
    </footer>
  );
};

export default Footer;
