import { useState } from "react";
import { Link } from "react-router-dom";
import { Menu, X } from "lucide-react";
import { Button } from "@/components/ui/button";
import AccountDropdown from "../ui/account-dropdown";
import { useAuth } from "@/hooks/use-auth";

const Navbar = () => {
  const [isOpen, setIsOpen] = useState(false);
  const { isAuthenticated } = useAuth();

  return (
    <nav className="bg-white shadow-md w-full z-10">
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div className="flex justify-between h-16">
          <div className="flex items-center">
            <Link to="/" className="flex-shrink-0 flex items-center">
              <span className="text-conecta-blue text-xl font-bold">
                Conecta<span className="text-conecta-green">Maré</span>
              </span>
            </Link>
            <div className="hidden md:ml-10 md:flex md:space-x-8">
              <Link
                to="/"
                className="text-gray-600 hover:text-conecta-blue px-3 py-2 text-sm font-medium"
              >
                Início
              </Link>
              <Link
                to="/professionals"
                className="text-gray-600 hover:text-conecta-blue px-3 py-2 text-sm font-medium"
              >
                Buscar Profissionais
              </Link>
              <Link
                to="/about"
                className="text-gray-600 hover:text-conecta-blue px-3 py-2 text-sm font-medium"
              >
                Como Funciona
              </Link>
              <Link
                to="/forpros"
                className="text-gray-600 hover:text-conecta-blue px-3 py-2 text-sm font-medium"
              >
                Para Profissionais
              </Link>
            </div>
          </div>
          <div className="hidden md:flex items-center">
            {isAuthenticated ? (
              <AccountDropdown />
            ) : (
              <div className="flex items-center space-x-4">
                <Link to="/login">
                  <Button
                    variant="outline"
                    className="border-conecta-blue text-conecta-blue hover:bg-conecta-blue hover:text-white"
                  >
                    Entrar
                  </Button>
                </Link>
                <Link to="/signup">
                  <Button className="bg-conecta-blue text-white hover:bg-conecta-blue-dark">
                    Cadastrar
                  </Button>
                </Link>
              </div>
            )}
          </div>
          <div className="flex md:hidden items-center">
            <button
              onClick={() => setIsOpen(!isOpen)}
              className="inline-flex items-center justify-center p-2 rounded-md text-gray-400 hover:text-gray-500 hover:bg-gray-100"
            >
              <span className="sr-only">Abrir menu</span>
              {isOpen ? <X size={24} /> : <Menu size={24} />}
            </button>
          </div>
        </div>
      </div>

      {/* Mobile menu */}
      {isOpen && (
        <div className="md:hidden animate-fade-in">
          <div className="px-2 pt-2 pb-3 space-y-1 sm:px-3 bg-white shadow-lg">
            <Link
              to="/"
              className="block px-3 py-2 text-gray-600 hover:text-conecta-blue hover:bg-gray-50 rounded-md"
            >
              Início
            </Link>
            <Link
              to="/professionals"
              className="block px-3 py-2 text-gray-600 hover:text-conecta-blue hover:bg-gray-50 rounded-md"
            >
              Buscar Profissionais
            </Link>
            <Link
              to="/about"
              className="block px-3 py-2 text-gray-600 hover:text-conecta-blue hover:bg-gray-50 rounded-md"
            >
              Como Funciona
            </Link>
            <Link
              to="/forpros"
              className="block px-3 py-2 text-gray-600 hover:text-conecta-blue hover:bg-gray-50 rounded-md"
            >
              Para Profissionais
            </Link>
            {isAuthenticated ? (
              <Link
                to="/dashboard"
                className="block px-3 py-2 text-conecta-blue font-medium hover:bg-gray-50 rounded-md"
              >
                Minha Conta
              </Link>
            ) : (
              <div className="space-y-2 pt-2">
                <Link to="/login" className="block w-full">
                  <Button
                    variant="outline"
                    className="w-full border-conecta-blue text-conecta-blue hover:bg-conecta-blue hover:text-white"
                  >
                    Entrar
                  </Button>
                </Link>
                <Link to="/signup" className="block w-full">
                  <Button className="w-full bg-conecta-blue text-white hover:bg-conecta-blue-dark">
                    Cadastrar
                  </Button>
                </Link>
              </div>
            )}
          </div>
        </div>
      )}
    </nav>
  );
};

export default Navbar;
