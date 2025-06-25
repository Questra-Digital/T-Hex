import {Montserrat, Outfit} from 'next/font/google';

export const montserrat = Montserrat({
    subsets: ['latin'],
    weight: ['700'], 
    style: ['normal'], 
    variable: '--font-montserrat',
  });

  export const outfit = Outfit({
    subsets: ['latin'],
    weight: ['400','500','700'],
    style: ['normal'],
    variable: '--font-outfit'
  })