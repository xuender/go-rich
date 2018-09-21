import { Component, OnInit, NgZone } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { Item } from '../api/item';
import { URL } from '../api/init'
import { ModalController, ActionSheetController, List } from '@ionic/angular';
import { ItemPage } from './item/item.page';
import { remove } from 'lodash'

@Component({
  selector: 'app-items',
  templateUrl: './items.page.html',
  styleUrls: ['./items.page.scss'],
})
export class ItemsPage implements OnInit {

  items$: Observable<Item[]>
  items: Item[] = []
  title = ''
  constructor(
    private http: HttpClient,
    private modalCtrl: ModalController,
    private ngZone: NgZone,
    private actionSheetCtrl: ActionSheetController
  ) {
    this.allItems()
  }
  allItems() {
    this.items$ = this.http.get<Item[]>(`${URL}/api/items`)
  }
  hasItem(items: Item[]): boolean {
    if (items) {
      this.items = items
      return items.length > 0
    }
    return false;
  }

  async create() {
    const modal = await this.modalCtrl.create({
      component: ItemPage,
      componentProps: {
        item: {
          name: '新商品',
        }
      }
    });
    modal.onDidDismiss().then(d => {
      if (d.data) {
        this.http.post<Item>(`${URL}/api/items`, d.data)
          .subscribe((i) => {
            this.items.push(i)
          })
      }
    })
    return await modal.present()
  }

  ngOnInit() {
  }
  async update(i: Item) {
    const modal = await this.modalCtrl.create({
      component: ItemPage,
      componentProps: {
        item: i,
      }
    });
    modal.onDidDismiss().then(d => {
      if (d.data) {
        this.http.put<Item>(`${URL}/api/items/${d.data.id}`, d.data)
          .subscribe()
      }
    })
    return await modal.present()
  }
  async del(i: Item, il: List) {
    const actionSheet = await this.actionSheetCtrl.create({
      header: `确认删除商品 [ ${i.name} ] ?`,
      buttons: [{
        text: '删除',
        role: 'destructive',
        icon: 'trash',
        handler: () => {
          il.closeSlidingItems()
          this.http.delete<Item>(`${URL}/api/items/${i.id}`)
            .subscribe(() => {
              remove(this.items, (d: Item) => i.id === d.id)
              this.ngZone.run(() => false)
            })
        }
      }, {
        text: '取消',
        icon: 'close',
        role: 'cancel',
      }]
    });
    await actionSheet.present();
  }
  search(event: CustomEvent) {
    const txt = event.detail.value
    this.title = `搜索: ${txt}`
    if (txt && txt.length > 1) {
      this.items.splice(0, this.items.length)
      this.items$ = this.http.get<Item[]>(`${URL}/api/items?search=${event.detail.value}`)
    }
  }
}
