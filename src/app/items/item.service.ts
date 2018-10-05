import { HttpClient } from '@angular/common/http';
import { Injectable, NgZone } from '@angular/core';

import { Item } from './item';
import { ObjService } from '../api/obj.service';
import { ModalController } from '@ionic/angular';
import { SelectPage } from './select/select.page';

@Injectable({
  providedIn: 'root'
})
export class ItemService extends ObjService<Item>{
  constructor(
    protected http: HttpClient,
    protected ngZone: NgZone,
    private modalCtrl: ModalController,
  ) {
    super(http, ngZone)
  }

  path() { return '/api/items' }

  async doSelect() {
    const modal = await this.modalCtrl.create({
      component: SelectPage,
      componentProps: { service: this }
    })
    await modal.present()
    return await modal.onDidDismiss()
  }
}
