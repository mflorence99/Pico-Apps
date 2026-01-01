# Pico Apps

Selection of apps to run on my Raspberry Pi Pico

## BME280

[Pinouts, Applications](https://electrocredible.com/bme280-pinout-specs-applications/)

[Pi Pico](https://electrocredible.com/raspberry-pi-pico-bme280-interfacing-guide-using-micropython/)

## How to Connect USB to Serial in WSL

> After all of this, neither `tinygo flash` nor `tinygo monitor` work properly under WSL. `/dev/ttyACM0` doesnt't appear to be stable.

Using PowerShell, Run as Administrator, and do this after every restart:

```sh
usbipd list
usbipd bind --busid 2-3
usbipd attach --wsl --busid 2-3
```

Now, run this in WSL and you should see the Raspberry Pi Pico:

```sh
lsusb
```

Run this and you should see /dev/ttyACM0:

```sh
ls -l /dev/tty*
```

And finally:

```sh
sudo chmod 666 /dev/ttyACM0
```

Some of this has to be repeated after you un/plug the USB cable to the Pico. 
