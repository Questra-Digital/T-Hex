import {Montserrat, Outfit, DM_Sans} from 'next/font/google';

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

  export const dm_sans = DM_Sans({
    subsets: ['latin'],
    weight: ['400','500','700'],
    style: ['normal'],
    variable: '--font-dmsans'
  })