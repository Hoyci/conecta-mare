
import { Suspense } from 'react';
import Navbar from '@/components/layout/Navbar';
import Footer from '@/components/layout/Footer';
import HeroSection from '@/components/home/HeroSection';
import FeaturesSection from '@/components/home/FeaturesSection';
import CategorySection from '@/components/home/CategorySection';
import HowItWorksSection from '@/components/home/HowItWorksSection';
import TestimonialsSection from '@/components/home/TestimonialsSection';
import CtaSection from '@/components/home/CtaSection';

const Index = () => {
  return (
    <div className="min-h-screen flex flex-col">
      <Navbar />
      <main className="flex-grow">
        <Suspense fallback={<div className="min-h-screen flex items-center justify-center">Carregando...</div>}>
          <HeroSection />
          <FeaturesSection />
          <CategorySection />
          <HowItWorksSection />
          <TestimonialsSection />
          <CtaSection />
        </Suspense>
      </main>
      <Footer />
    </div>
  );
};

export default Index;
