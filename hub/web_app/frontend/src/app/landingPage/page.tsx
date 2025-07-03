import HeroSection from "@/components/HeroSection";
import FeaturesSection from "@/components/FeaturesSection";
import PricingSection from "@/components/PricingSection";
import IntegrationsSection from "@/components/IntegrationsSection";
import WhatItDo from "@/components/WhatItDoSection";
import Footer from "@/components/Footer";
export default function Page() {
  return (
    <>
      <HeroSection />
      <WhatItDo />
      <FeaturesSection />
      <IntegrationsSection />
      <PricingSection />
      <Footer />
    </>
  );
}