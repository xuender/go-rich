import { Component, OnInit } from '@angular/core';
import { Platform } from '@ionic/angular';

@Component({
  selector: 'app-tabs',
  templateUrl: 'tabs.page.html',
  styleUrls: ['tabs.page.scss']
})
export class TabsPage implements OnInit {
  constructor(
    private plt: Platform
  ) {
  }
  ngOnInit() {
    if (!this.plt.is('mobile') && !this.plt.is('mobileweb')) {
      this.selectQr()
    }
  }
  async selectQr() {
    console.log('qr...')
  }
}
