import { ModalController, List, ActionSheetController } from '@ionic/angular';

import { Obj } from './obj';
import { ObjService } from './obj.service';

export abstract class ObjsPage<T extends Obj>{
  constructor(
    protected modalCtrl: ModalController,
    protected actionSheetCtrl: ActionSheetController,
  ) { }

  abstract get service(): ObjService<T>

  abstract get page(): any

  abstract get newObj(): any

  abstract get title(): string

  cancel() {
    this.modalCtrl.dismiss();
  }

  async create() {
    const modal = await this.modalCtrl.create({
      component: this.page,
      componentProps: { obj: this.newObj }
    })
    return await modal.present()
  }

  async update(t: T) {
    const modal = await this.modalCtrl.create({
      component: this.page,
      componentProps: { obj: Object.assign({}, t) },
    });
    return await modal.present()
  }

  async del(t: T, il: List) {
    const actionSheet = await this.actionSheetCtrl.create({
      header: `确认删除${this.title} [ ${t.name} ] ?`,
      buttons: [{
        text: '删除',
        role: 'destructive',
        icon: 'trash',
        handler: () => {
          this.service.delete(t)
        }
      }, {
        text: '取消',
        icon: 'close',
        role: 'cancel',
      }]
    });
    await actionSheet.present();
    await il.closeSlidingItems()
  }
}
