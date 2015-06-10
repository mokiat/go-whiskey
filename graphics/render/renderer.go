package render

import "github.com/momchil-atanasov/go-whiskey/graphics/client"

//go:generate counterfeiter -o render_fakes/fake_renderer.go ./ Renderer

type Renderer interface {
	Render(Queue)
}

func NewRenderer(graphicsClient client.GraphicsClient) Renderer {
	return &renderer{
		graphicsClient: graphicsClient,
	}
}

type renderer struct {
	graphicsClient client.GraphicsClient
}

func (r *renderer) Render(queue Queue) {
	for item, has := queue.Pop(); has; item, has = queue.Pop() {
		r.graphicsClient.UseProgram(item.ProgramId)

		for _, attribute := range item.Attributes {
			r.graphicsClient.EnableAttribute(attribute.Location)
			r.graphicsClient.BindVertexBuffer(attribute.VertexBufferId)
			r.graphicsClient.ConfigureAttribute(attribute.Location, attribute.Components, attribute.StrideInBytes, attribute.OffsetInBytes)
		}

		for _, uniform := range item.Uniforms {
			switch uniform.Type {
			case Vec4UniformType:
				r.graphicsClient.SetVec4Uniform(uniform.Location, uniform.Values)
			case Mat4x4UniformType:
				r.graphicsClient.SetMat4x4Uniform(uniform.Location, uniform.Values)
			}
		}

		channel := 0
		for _, texture := range item.Textures {
			switch texture.Type {
			case TwoDTextureType:
				r.graphicsClient.Bind2DTexture(channel, texture.TextureId)
			case CubeTextureType:
				r.graphicsClient.BindCubeTexture(channel, texture.TextureId)
			}
			r.graphicsClient.SetSamplerUniform(texture.Location, channel)
			channel++
		}

		for _, element := range item.Elements {
			r.graphicsClient.BindIndexBuffer(element.IndexBufferId)
			switch element.Type {
			case TrianglesElementType:
				r.graphicsClient.DrawTriangles(element.IndexCount, element.IndexOffsetInBytes)
			case LinesElementType:
				r.graphicsClient.DrawLines(element.IndexCount, element.IndexOffsetInBytes)
			}
		}

		for _, attribute := range item.Attributes {
			r.graphicsClient.DisableAttribute(attribute.Location)
		}
	}
}
