# gbprinter
Convert images to GameBoy screenshots.

Installation
---

```bash
go get github.com/ProfOak/gbprinter
```

Usage
---

```bash
gbprinter image.png
```

The resulting image will be named `image_color-palette.png`


Extra
---

```bash
gbprinter -help
Usage of gbprinter:
  -palette string
          Color choices: down, downa, downb, grayscale, greenscale, left, lefta, leftb, right, righta, rightb, up, upa, upb (default "grayscale")
```

Examples
---

### Grayscales

![original](images/amarao.jpg)
![grayscale](images/amarao_grayscale.png)

![original](images/catac.jpg)
![grayscale](images/catac_grayscale.png)

![original](images/integrating_pistol.jpg)
![grayscale](images/integrating_pistol_grayscale.png)

![original](images/kirby.png)
![grayscale](images/kirby_grayscale.png)

![original](images/wizard_cat.png)
![grayscale](images/wizard_cat_grayscale.png)

### Greenscale and the other color palettes

![original](images/skate_or_die.jpg)
![greenscale](images/skate_or_die_greenscale.png)
![all of the color palettes](images/skate_or_die_collage.jpg)
