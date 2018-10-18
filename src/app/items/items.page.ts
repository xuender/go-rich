import { Component, EventEmitter } from '@angular/core';
import { ModalController, ActionSheetController } from '@ionic/angular';

import { Item } from './item';
import { Paging } from '../api/paging';
import { ItemPage } from './item/item.page';
import { ItemService } from './item.service';
import { TagService } from '../tags/tag.service';
import { ObjsPlusPage } from '../api/objs.puls.page';
import { UploadPage } from '../setting/xlsxes/upload/upload.page';

@Component({
  selector: 'app-items',
  templateUrl: './items.page.html',
  styleUrls: ['./items.page.scss'],
})
export class ItemsPage extends ObjsPlusPage<Item> {
  tags: string[] = []
  paging: Paging<Item> = { data: [], total: 0 }
  private reload: EventEmitter<boolean>
  constructor(
    public itemService: ItemService,
    protected tagService: TagService,
    protected modalCtrl: ModalController,
    protected actionSheetCtrl: ActionSheetController,
  ) {
    super(tagService, modalCtrl, actionSheetCtrl)
    this.reload = new EventEmitter()
    this.reload.subscribe(() => this.reset())
    this.itemService.itemsReload = this.reload
  }

  get tagKey() { return 'tag-I' };
  get service() { return this.itemService }
  get page() { return ItemPage }
  get newObj() { return { name: '', price: 0, extend: {}, tags: [] } }
  get title() { return '商品' }

  async upload() {
    const modal = await this.modalCtrl.create({
      component: UploadPage,
      componentProps: { url: 'items/file' }
    });
    await modal.present()
    await modal.onDidDismiss()
    await this.reset()
  }
}
