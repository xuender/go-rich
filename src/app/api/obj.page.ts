import { ModalController, NavParams } from '@ionic/angular';
import { Obj } from './obj';
import { ObjService } from './obj.service';

export abstract class ObjPage<T extends Obj> {
  obj: T
  constructor(
    protected modalCtrl: ModalController,
    protected navParams: NavParams
  ) {
    this.obj = this.navParams.get('obj')
  }

  abstract get service(): ObjService<T>

  cancel() {
    this.modalCtrl.dismiss();
  }

  async save() {
    await this.modalCtrl.dismiss(await this.service.save(this.obj))
  }
}
