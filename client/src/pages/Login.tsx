import Navbar from "@/components/layout/Navbar";
import Footer from "@/components/layout/Footer";
import LoginForm from "@/components/auth/LoginForm";

const Login = () => {
  return (
    <div className="min-h-screen flex flex-col">
      <Navbar />
      <main className="flex-grow flex items-center justify-center bg-gray-50 py-12">
        <div className="w-full max-w-md p-6 bg-white rounded-lg shadow-md">
          <LoginForm />
        </div>
      </main>
      <Footer />
    </div>
  );
};

export default Login;
