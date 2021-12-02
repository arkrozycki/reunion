import {
  createApp as createClientApp,
  h,
  VNode,
  resolveDynamicComponent,
  Transition,
} from "vue";
import { RouterView } from "vue-router";

export async function createApp() {
  const app = createClientApp({
    // This is the global app setup function
    setup() {
      return () => {
        const defaultSlot = ({ Component: _Component }: any) => {
          const Component = resolveDynamicComponent(_Component) as VNode;

          return [
            h(
              Transition,
              { name: "fade-slow", mode: "out-in" },
              {
                default: () => [h(Component)],
              }
            ),
          ];
        };
        return h(RouterView, null, {
          default: defaultSlot,
        });
      };
    },
  });

  return app;
}
