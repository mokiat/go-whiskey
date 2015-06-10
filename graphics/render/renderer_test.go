package render_test

import (
	"github.com/momchil-atanasov/go-whiskey/graphics/client"
	"github.com/momchil-atanasov/go-whiskey/graphics/client/client_fakes"
	. "github.com/momchil-atanasov/go-whiskey/graphics/render"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Renderer", func() {
	var graphicsClient *client_fakes.FakeGraphicsClient
	var queue Queue
	var renderer Renderer

	BeforeEach(func() {
		graphicsClient = new(client_fakes.FakeGraphicsClient)
		queue = NewFIFOQueue()
		renderer = NewRenderer(graphicsClient)
	})

	It("should not be nil", func() {
		Ω(renderer).ShouldNot(BeNil())
	})

	Describe("Render flow", func() {
		var meshShaderProgramId client.ProgramId
		var skyboxShaderProgramId client.ProgramId

		var coordLocation client.AttributeLocation
		var coordBuffer client.BufferId
		var coordComponents int
		var coordStride int
		var coordOffset int

		var texCoordLocation client.AttributeLocation
		var texCoordBuffer client.BufferId
		var texCoordComponents int
		var texCoordStride int
		var texCoordOffset int

		var projectionMatrixLocation client.UniformLocation
		var projectionMatrixType UniformType
		var projectionMatrixData []float32

		var diffuseColorLocation client.UniformLocation
		var diffuseColorType UniformType
		var diffuseColorData []float32

		var diffuseTextureLocation client.UniformLocation
		var diffuseTextureType TextureType
		var diffuseTextureId client.TextureId

		var skyboxTextureLocation client.UniformLocation
		var skyboxTextureType TextureType
		var skyboxTextureId client.TextureId

		var airplaneBufferId client.BufferId
		var airplaneType ElementType
		var airplaneIndexCount int
		var airplaneIndexOffset int

		var shuttleBufferId client.BufferId
		var shuttleType ElementType
		var shuttleIndexCount int
		var shuttleIndexOffset int

		BeforeEach(func() {
			meshShaderProgramId = client.ProgramId(10)
			skyboxShaderProgramId = client.ProgramId(16)

			coordLocation = client.AttributeLocation(0)
			coordBuffer = client.BufferId(2)
			coordComponents = 3
			coordStride = 3 * 4
			coordOffset = 5 * 4

			texCoordLocation = client.AttributeLocation(1)
			texCoordBuffer = client.BufferId(3)
			texCoordComponents = 2
			texCoordStride = 2 * 4
			texCoordOffset = 8 * 4

			projectionMatrixLocation = client.UniformLocation(2)
			projectionMatrixType = Mat4x4UniformType
			projectionMatrixData = []float32{0.3, 2.1, 4.3, 2.9}

			diffuseColorLocation = client.UniformLocation(3)
			diffuseColorType = Vec4UniformType
			diffuseColorData = []float32{0.2, 1.3}

			diffuseTextureLocation = client.UniformLocation(3)
			diffuseTextureType = TwoDTextureType
			diffuseTextureId = client.TextureId(13)

			skyboxTextureLocation = client.UniformLocation(4)
			skyboxTextureType = CubeTextureType
			skyboxTextureId = client.TextureId(14)

			airplaneBufferId = client.BufferId(5)
			airplaneType = TrianglesElementType
			airplaneIndexCount = 10
			airplaneIndexOffset = 52

			shuttleBufferId = client.BufferId(6)
			shuttleType = LinesElementType
			shuttleIndexCount = 20
			shuttleIndexOffset = 102
		})

		JustBeforeEach(func() {
			renderer.Render(queue)
		})

		Context("when a single item is rendered", func() {
			BeforeEach(func() {
				queue.Push(Item{
					ProgramId: meshShaderProgramId,
					Attributes: []AttributeEntry{
						AttributeEntry{
							Location:       coordLocation,
							VertexBufferId: coordBuffer,
							Components:     coordComponents,
							StrideInBytes:  coordStride,
							OffsetInBytes:  coordOffset,
						},
					},
					Uniforms: []UniformEntry{
						UniformEntry{
							Location: projectionMatrixLocation,
							Type:     projectionMatrixType,
							Values:   projectionMatrixData,
						},
					},
					Textures: []TextureEntry{
						TextureEntry{
							Location:  diffuseTextureLocation,
							Type:      diffuseTextureType,
							TextureId: diffuseTextureId,
						},
					},
					Elements: []ElementEntry{
						ElementEntry{
							IndexBufferId:      airplaneBufferId,
							Type:               airplaneType,
							IndexCount:         airplaneIndexCount,
							IndexOffsetInBytes: airplaneIndexOffset,
						},
					},
				})
			})

			It("should have rendered the item", func() {
				By("enabling the shader program")
				Ω(graphicsClient.UseProgramCallCount()).Should(Equal(1))
				argProgramId := graphicsClient.UseProgramArgsForCall(0)
				Ω(argProgramId).Should(Equal(meshShaderProgramId))

				By("enabling all the attributes")
				Ω(graphicsClient.EnableAttributeCallCount()).Should(Equal(1))
				argLocation := graphicsClient.EnableAttributeArgsForCall(0)
				Ω(argLocation).Should(Equal(coordLocation))

				By("configuring all the attributes")
				Ω(graphicsClient.BindVertexBufferCallCount()).Should(Equal(1))
				argBufferId := graphicsClient.BindVertexBufferArgsForCall(0)
				Ω(argBufferId).Should(Equal(coordBuffer))

				Ω(graphicsClient.ConfigureAttributeCallCount()).Should(Equal(1))
				argLocation, argComponents, argStride, argOffset := graphicsClient.ConfigureAttributeArgsForCall(0)
				Ω(argLocation).Should(Equal(coordLocation))
				Ω(argComponents).Should(Equal(coordComponents))
				Ω(argStride).Should(Equal(coordStride))
				Ω(argOffset).Should(Equal(coordOffset))

				By("configuring all the uniforms")
				Ω(graphicsClient.SetMat4x4UniformCallCount()).Should(Equal(1))
				argLocation, argData := graphicsClient.SetMat4x4UniformArgsForCall(0)
				Ω(argLocation).Should(Equal(projectionMatrixLocation))
				Ω(argData).Should(Equal(projectionMatrixData))

				By("configuring all the texture uniforms")
				Ω(graphicsClient.Bind2DTextureCallCount()).Should(Equal(1))
				argChannel, argTextureId := graphicsClient.Bind2DTextureArgsForCall(0)
				Ω(argChannel).Should(Equal(0))
				Ω(argTextureId).Should(Equal(diffuseTextureId))

				Ω(graphicsClient.SetSamplerUniformCallCount()).Should(Equal(1))
				argLocation, argChannel = graphicsClient.SetSamplerUniformArgsForCall(0)
				Ω(argLocation).Should(Equal(diffuseTextureLocation))
				Ω(argChannel).Should(Equal(0))

				By("rendering all the elements")
				Ω(graphicsClient.BindIndexBufferCallCount()).Should(Equal(1))
				argBufferId = graphicsClient.BindIndexBufferArgsForCall(0)
				Ω(argBufferId).Should(Equal(airplaneBufferId))

				Ω(graphicsClient.DrawTrianglesCallCount()).Should(Equal(1))
				argIndexCount, argIndexOffset := graphicsClient.DrawTrianglesArgsForCall(0)
				Ω(argIndexCount).Should(Equal(airplaneIndexCount))
				Ω(argIndexOffset).Should(Equal(airplaneIndexOffset))

				By("disabling all the attributes")
				Ω(graphicsClient.DisableAttributeCallCount()).Should(Equal(1))
				argLocation = graphicsClient.DisableAttributeArgsForCall(0)
				Ω(argLocation).Should(Equal(coordLocation))
			})
		})

		Context("when a multiple items are rendered", func() {
			BeforeEach(func() {
				queue.Push(Item{
					ProgramId: meshShaderProgramId,
					Attributes: []AttributeEntry{
						AttributeEntry{
							Location:       coordLocation,
							VertexBufferId: coordBuffer,
							Components:     coordComponents,
							StrideInBytes:  coordStride,
							OffsetInBytes:  coordOffset,
						},
					},
					Uniforms: []UniformEntry{
						UniformEntry{
							Location: projectionMatrixLocation,
							Type:     projectionMatrixType,
							Values:   projectionMatrixData,
						},
					},
					Textures: []TextureEntry{
						TextureEntry{
							Location:  diffuseTextureLocation,
							Type:      diffuseTextureType,
							TextureId: diffuseTextureId,
						},
					},
					Elements: []ElementEntry{
						ElementEntry{
							IndexBufferId:      airplaneBufferId,
							Type:               airplaneType,
							IndexCount:         airplaneIndexCount,
							IndexOffsetInBytes: airplaneIndexOffset,
						},
					},
				})
				queue.Push(Item{
					ProgramId: skyboxShaderProgramId,
					Attributes: []AttributeEntry{
						AttributeEntry{
							Location:       texCoordLocation,
							VertexBufferId: texCoordBuffer,
							Components:     texCoordComponents,
							StrideInBytes:  texCoordStride,
							OffsetInBytes:  texCoordOffset,
						},
					},
					Uniforms: []UniformEntry{
						UniformEntry{
							Location: diffuseColorLocation,
							Type:     diffuseColorType,
							Values:   diffuseColorData,
						},
					},
					Textures: []TextureEntry{
						TextureEntry{
							Location:  skyboxTextureLocation,
							Type:      skyboxTextureType,
							TextureId: skyboxTextureId,
						},
					},
					Elements: []ElementEntry{
						ElementEntry{
							IndexBufferId:      shuttleBufferId,
							Type:               shuttleType,
							IndexCount:         shuttleIndexCount,
							IndexOffsetInBytes: shuttleIndexOffset,
						},
					},
				})
			})

			It("should have rendered the first item", func() {
				By("enabling the shader program")
				argProgramId := graphicsClient.UseProgramArgsForCall(0)
				Ω(argProgramId).Should(Equal(meshShaderProgramId))

				By("enabling all the attributes")
				argLocation := graphicsClient.EnableAttributeArgsForCall(0)
				Ω(argLocation).Should(Equal(coordLocation))

				By("configuring all the attributes")
				argBufferId := graphicsClient.BindVertexBufferArgsForCall(0)
				Ω(argBufferId).Should(Equal(coordBuffer))

				argLocation, argComponents, argStride, argOffset := graphicsClient.ConfigureAttributeArgsForCall(0)
				Ω(argLocation).Should(Equal(coordLocation))
				Ω(argComponents).Should(Equal(coordComponents))
				Ω(argStride).Should(Equal(coordStride))
				Ω(argOffset).Should(Equal(coordOffset))

				By("configuring all the uniforms")
				argLocation, argData := graphicsClient.SetMat4x4UniformArgsForCall(0)
				Ω(argLocation).Should(Equal(projectionMatrixLocation))
				Ω(argData).Should(Equal(projectionMatrixData))

				By("configuring all the texture uniforms")
				argChannel, argTextureId := graphicsClient.Bind2DTextureArgsForCall(0)
				Ω(argChannel).Should(Equal(0))
				Ω(argTextureId).Should(Equal(diffuseTextureId))

				argLocation, argChannel = graphicsClient.SetSamplerUniformArgsForCall(0)
				Ω(argLocation).Should(Equal(diffuseTextureLocation))
				Ω(argChannel).Should(Equal(0))

				By("rendering all the elements")
				argBufferId = graphicsClient.BindIndexBufferArgsForCall(0)
				Ω(argBufferId).Should(Equal(airplaneBufferId))

				argIndexCount, argIndexOffset := graphicsClient.DrawTrianglesArgsForCall(0)
				Ω(argIndexCount).Should(Equal(airplaneIndexCount))
				Ω(argIndexOffset).Should(Equal(airplaneIndexOffset))

				By("disabling all the attributes")
				argLocation = graphicsClient.DisableAttributeArgsForCall(0)
				Ω(argLocation).Should(Equal(coordLocation))
			})

			It("should have rendered the second item", func() {
				By("enabling the shader programs")
				argProgramId := graphicsClient.UseProgramArgsForCall(1)
				Ω(argProgramId).Should(Equal(skyboxShaderProgramId))

				By("enabling all the attributes")
				argLocation := graphicsClient.EnableAttributeArgsForCall(1)
				Ω(argLocation).Should(Equal(texCoordLocation))

				By("configuring all the attributes")
				argBufferId := graphicsClient.BindVertexBufferArgsForCall(1)
				Ω(argBufferId).Should(Equal(texCoordBuffer))

				argLocation, argComponents, argStride, argOffset := graphicsClient.ConfigureAttributeArgsForCall(1)
				Ω(argLocation).Should(Equal(texCoordLocation))
				Ω(argComponents).Should(Equal(texCoordComponents))
				Ω(argStride).Should(Equal(texCoordStride))
				Ω(argOffset).Should(Equal(texCoordOffset))

				By("configuring all the uniforms")
				argLocation, argData := graphicsClient.SetVec4UniformArgsForCall(0)
				Ω(argLocation).Should(Equal(diffuseColorLocation))
				Ω(argData).Should(Equal(diffuseColorData))

				By("configuring all the texture uniforms")
				argChannel, argTextureId := graphicsClient.BindCubeTextureArgsForCall(0)
				Ω(argChannel).Should(Equal(0))
				Ω(argTextureId).Should(Equal(skyboxTextureId))

				argLocation, argChannel = graphicsClient.SetSamplerUniformArgsForCall(1)
				Ω(argLocation).Should(Equal(skyboxTextureLocation))
				Ω(argChannel).Should(Equal(0))

				By("rendering all the elements")
				argBufferId = graphicsClient.BindIndexBufferArgsForCall(1)
				Ω(argBufferId).Should(Equal(shuttleBufferId))

				argIndexCount, argIndexOffset := graphicsClient.DrawLinesArgsForCall(0)
				Ω(argIndexCount).Should(Equal(shuttleIndexCount))
				Ω(argIndexOffset).Should(Equal(shuttleIndexOffset))

				By("disabling all the attributes")
				argLocation = graphicsClient.DisableAttributeArgsForCall(1)
				Ω(argLocation).Should(Equal(texCoordLocation))
			})
		})

		Context("when multiple attributes are assigned", func() {
			BeforeEach(func() {
				queue.Push(Item{
					ProgramId: meshShaderProgramId,
					Attributes: []AttributeEntry{
						AttributeEntry{
							Location:       coordLocation,
							VertexBufferId: coordBuffer,
							Components:     coordComponents,
							StrideInBytes:  coordStride,
							OffsetInBytes:  coordOffset,
						},
						AttributeEntry{
							Location:       texCoordLocation,
							VertexBufferId: texCoordBuffer,
							Components:     texCoordComponents,
							StrideInBytes:  texCoordStride,
							OffsetInBytes:  texCoordOffset,
						},
					},
					Uniforms: []UniformEntry{},
					Textures: []TextureEntry{},
					Elements: []ElementEntry{},
				})
			})

			It("should have enabled all attributes", func() {
				Ω(graphicsClient.EnableAttributeCallCount()).Should(Equal(2))
				argLocation := graphicsClient.EnableAttributeArgsForCall(0)
				Ω(argLocation).Should(Equal(coordLocation))
				argLocation = graphicsClient.EnableAttributeArgsForCall(1)
				Ω(argLocation).Should(Equal(texCoordLocation))
			})

			It("should have configured all attributes", func() {
				Ω(graphicsClient.BindVertexBufferCallCount()).Should(Equal(2))
				Ω(graphicsClient.ConfigureAttributeCallCount()).Should(Equal(2))

				argBufferId := graphicsClient.BindVertexBufferArgsForCall(0)
				Ω(argBufferId).Should(Equal(coordBuffer))

				argLocation, argComponents, argStride, argOffset := graphicsClient.ConfigureAttributeArgsForCall(0)
				Ω(argLocation).Should(Equal(coordLocation))
				Ω(argComponents).Should(Equal(coordComponents))
				Ω(argStride).Should(Equal(coordStride))
				Ω(argOffset).Should(Equal(coordOffset))

				argBufferId = graphicsClient.BindVertexBufferArgsForCall(1)
				Ω(argBufferId).Should(Equal(texCoordBuffer))

				argLocation, argComponents, argStride, argOffset = graphicsClient.ConfigureAttributeArgsForCall(1)
				Ω(argLocation).Should(Equal(texCoordLocation))
				Ω(argComponents).Should(Equal(texCoordComponents))
				Ω(argStride).Should(Equal(texCoordStride))
				Ω(argOffset).Should(Equal(texCoordOffset))
			})

			It("should have disabled all attributes", func() {
				Ω(graphicsClient.DisableAttributeCallCount()).Should(Equal(2))
				argLocation := graphicsClient.DisableAttributeArgsForCall(0)
				Ω(argLocation).Should(Equal(coordLocation))
				argLocation = graphicsClient.DisableAttributeArgsForCall(1)
				Ω(argLocation).Should(Equal(texCoordLocation))
			})
		})

		Context("when multiple uniforms are assigned", func() {
			BeforeEach(func() {
				queue.Push(Item{
					ProgramId:  meshShaderProgramId,
					Attributes: []AttributeEntry{},
					Uniforms: []UniformEntry{
						UniformEntry{
							Location: projectionMatrixLocation,
							Type:     projectionMatrixType,
							Values:   projectionMatrixData,
						},
						UniformEntry{
							Location: diffuseColorLocation,
							Type:     diffuseColorType,
							Values:   diffuseColorData,
						},
					},
					Textures: []TextureEntry{},
					Elements: []ElementEntry{},
				})
			})

			It("should have configured all uniforms", func() {
				Ω(graphicsClient.SetMat4x4UniformCallCount()).Should(Equal(1))
				argLocation, argData := graphicsClient.SetMat4x4UniformArgsForCall(0)
				Ω(argLocation).Should(Equal(projectionMatrixLocation))
				Ω(argData).Should(Equal(projectionMatrixData))

				Ω(graphicsClient.SetVec4UniformCallCount()).Should(Equal(1))
				argLocation, argData = graphicsClient.SetVec4UniformArgsForCall(0)
				Ω(argLocation).Should(Equal(diffuseColorLocation))
				Ω(argData).Should(Equal(diffuseColorData))
			})
		})

		Context("when multiple textures are assigned", func() {
			BeforeEach(func() {
				queue.Push(Item{
					ProgramId:  meshShaderProgramId,
					Attributes: []AttributeEntry{},
					Uniforms:   []UniformEntry{},
					Textures: []TextureEntry{
						TextureEntry{
							Location:  diffuseTextureLocation,
							Type:      diffuseTextureType,
							TextureId: diffuseTextureId,
						},
						TextureEntry{
							Location:  skyboxTextureLocation,
							Type:      skyboxTextureType,
							TextureId: skyboxTextureId,
						},
					},
					Elements: []ElementEntry{},
				})
			})

			It("should have configured all texture uniforms", func() {
				Ω(graphicsClient.Bind2DTextureCallCount()).Should(Equal(1))
				Ω(graphicsClient.BindCubeTextureCallCount()).Should(Equal(1))
				Ω(graphicsClient.SetSamplerUniformCallCount()).Should(Equal(2))

				argChannel, argTextureId := graphicsClient.Bind2DTextureArgsForCall(0)
				Ω(argChannel).Should(Equal(0))
				Ω(argTextureId).Should(Equal(diffuseTextureId))

				argLocation, argChannel := graphicsClient.SetSamplerUniformArgsForCall(0)
				Ω(argLocation).Should(Equal(diffuseTextureLocation))
				Ω(argChannel).Should(Equal(0))

				argChannel, argTextureId = graphicsClient.BindCubeTextureArgsForCall(0)
				Ω(argChannel).Should(Equal(1))
				Ω(argTextureId).Should(Equal(skyboxTextureId))

				argLocation, argChannel = graphicsClient.SetSamplerUniformArgsForCall(1)
				Ω(argLocation).Should(Equal(skyboxTextureLocation))
				Ω(argChannel).Should(Equal(1))
			})
		})

		Context("when multiple elements are configured", func() {
			BeforeEach(func() {
				queue.Push(Item{
					ProgramId:  meshShaderProgramId,
					Attributes: []AttributeEntry{},
					Uniforms:   []UniformEntry{},
					Textures:   []TextureEntry{},
					Elements: []ElementEntry{
						ElementEntry{
							IndexBufferId:      airplaneBufferId,
							Type:               airplaneType,
							IndexCount:         airplaneIndexCount,
							IndexOffsetInBytes: airplaneIndexOffset,
						},
						ElementEntry{
							IndexBufferId:      shuttleBufferId,
							Type:               shuttleType,
							IndexCount:         shuttleIndexCount,
							IndexOffsetInBytes: shuttleIndexOffset,
						},
					},
				})
			})

			It("should have rendered all elements", func() {
				Ω(graphicsClient.BindIndexBufferCallCount()).Should(Equal(2))
				Ω(graphicsClient.DrawTrianglesCallCount()).Should(Equal(1))
				Ω(graphicsClient.DrawLinesCallCount()).Should(Equal(1))

				argBufferId := graphicsClient.BindIndexBufferArgsForCall(0)
				Ω(argBufferId).Should(Equal(airplaneBufferId))

				argIndexCount, argIndexOffset := graphicsClient.DrawTrianglesArgsForCall(0)
				Ω(argIndexCount).Should(Equal(airplaneIndexCount))
				Ω(argIndexOffset).Should(Equal(airplaneIndexOffset))

				argBufferId = graphicsClient.BindIndexBufferArgsForCall(1)
				Ω(argBufferId).Should(Equal(shuttleBufferId))

				argIndexCount, argIndexOffset = graphicsClient.DrawLinesArgsForCall(0)
				Ω(argIndexCount).Should(Equal(shuttleIndexCount))
				Ω(argIndexOffset).Should(Equal(shuttleIndexOffset))
			})
		})
	})
})
